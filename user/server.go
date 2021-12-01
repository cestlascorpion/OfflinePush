package user

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	pb "github.com/cestlascorpion/offlinepush/proto"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	*pb.UnimplementedUserServer
	mgr  *Mgr
	auth *core.AuthCache
}

func NewServer(conf *core.PushConfig) (*Server, error) {
	mgr, err := NewUserMgr()
	if err != nil {
		log.Errorf("new user mgr err %+v", err)
		return nil, err
	}

	agt, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		log.Errorf("new getui agent err %+v", err)
		return nil, err
	}
	err = mgr.RegisterAgent(core.UniqueId{PushAgent: conf.TestApp.PushAgent, BundleId: conf.TestApp.BundleId}, agt)
	if err != nil {
		log.Errorf("register getui agent err %+v", err)
		return nil, err
	}

	auth, err := core.NewAuthCache()
	if err != nil {
		log.Errorf("new auth cache err %+v", err)
		return nil, err
	}
	err = auth.Start(context.Background())
	if err != nil {
		log.Errorf("start auth cache err %+v", err)
		return nil, err
	}

	return &Server{
		mgr:  mgr,
		auth: auth,
	}, nil
}

func (s *Server) BindAlias(ctx context.Context, in *pb.BindAliasReq) (*pb.BindAliasResp, error) {
	out := &pb.BindAliasResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.DataList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	aliasList := &AliasList{
		DataList: make([]*DataList, 0),
	}
	for i := range in.DataList {
		aliasList.DataList = append(aliasList.DataList, &DataList{
			Cid:   in.DataList[i].Cid,
			Alias: in.DataList[i].Alias,
		})
	}

	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.BindAlias(uniqueId, aliasList, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			err = s.mgr.BindAlias(uniqueId, aliasList, auth.Token)
		}
		if err != nil {
			log.Errorf("bind alias err %+v", err)
			return out, err
		}
	}
	return out, nil
}

func (s *Server) QueryAliasByCid(ctx context.Context, in *pb.QueryAliasReq) (*pb.QueryAliasResp, error) {
	out := &pb.QueryAliasResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	alias, err := s.mgr.QueryAliasByCid(uniqueId, in.CId, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			alias, err = s.mgr.QueryAliasByCid(uniqueId, in.CId, auth.Token)
		}
		if err != nil && err.Error() != core.InvalidTargetErr {
			log.Errorf("query alias err %+v", err)
			return out, err
		}
	}
	out.Alias = alias
	return out, nil
}

func (s *Server) QueryCidByAlias(ctx context.Context, in *pb.QueryCidReq) (*pb.QueryCidResp, error) {
	out := &pb.QueryCidResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.Alias) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	cidList, err := s.mgr.QueryCidByAlias(uniqueId, in.Alias, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			cidList, err = s.mgr.QueryCidByAlias(uniqueId, in.Alias, auth.Token)
		}
		if err != nil && err.Error() != core.InvalidTargetErr {
			log.Errorf("query cid err %+v", err)
			return out, err
		}
	}
	out.CIdList = cidList
	return out, nil
}

func (s *Server) UnbindAlias(ctx context.Context, in *pb.UnbindAliasReq) (*pb.UnbindAliasResp, error) {
	out := &pb.UnbindAliasResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.DataList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	aliasList := &AliasList{
		DataList: make([]*DataList, 0),
	}
	for i := range in.DataList {
		aliasList.DataList = append(aliasList.DataList, &DataList{
			Cid:   in.DataList[i].Cid,
			Alias: in.DataList[i].Alias,
		})
	}

	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.UnbindAlias(uniqueId, aliasList, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			err = s.mgr.UnbindAlias(uniqueId, aliasList, auth.Token)
		}
		if err != nil {
			log.Errorf("unbind alias err %+v", err)
			return out, err
		}
	}
	return out, nil
}

func (s *Server) RevokeAlias(ctx context.Context, in *pb.RevokeAliasReq) (*pb.RevokeAliasResp, error) {
	out := &pb.RevokeAliasResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.Alias) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.RevokeAlias(uniqueId, in.Alias, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			err = s.mgr.RevokeAlias(uniqueId, in.Alias, auth.Token)
		}
		if err != nil {
			log.Errorf("revoke alias err %+v", err)
			return out, err
		}
	}
	return out, nil
}

func (s *Server) BindUserWithTag(ctx context.Context, in *pb.BindUserWithTagReq) (*pb.BindUserWithTagResp, error) {
	out := &pb.BindUserWithTagResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CId) == 0 || len(in.TagList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	tagList := &CustomTagList{
		TagList: in.TagList,
	}

	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.BindUserWithTag(uniqueId, in.CId, tagList, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			err = s.mgr.BindUserWithTag(uniqueId, in.CId, tagList, auth.Token)
		}
		if err != nil {
			log.Errorf("bind user with tag err %+v", err)
			return out, err
		}
	}
	return out, nil
}

func (s *Server) BindTagWithUser(ctx context.Context, in *pb.BindTagWithUserReq) (*pb.BindTagWithUserResp, error) {
	out := &pb.BindTagWithUserResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.Tag) == 0 || len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	cidList := &CidList{
		CidList: in.CIdList,
	}

	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.mgr.BindTagWithUser(uniqueId, in.Tag, cidList, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			resp, err = s.mgr.BindTagWithUser(uniqueId, in.Tag, cidList, auth.Token)
		}
		if err != nil {
			log.Errorf("bind tag with user err %+v", err)
			return out, err
		}
	}
	out.ResultList = make([]*pb.BindTagWithUserResp_Result, 0)
	for cid, success := range resp {
		out.ResultList = append(out.ResultList, &pb.BindTagWithUserResp_Result{
			CId:     cid,
			Success: success,
		})
	}
	return out, nil
}

func (s *Server) UnbindTagFromUser(ctx context.Context, in *pb.UnbindTagFromUserReq) (*pb.UnbindTagFromUserResp, error) {
	out := &pb.UnbindTagFromUserResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.Tag) == 0 || len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	cidList := &CidList{
		CidList: in.CIdList,
	}

	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.mgr.UnbindTagFromUser(uniqueId, in.Tag, cidList, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			resp, err = s.mgr.UnbindTagFromUser(uniqueId, in.Tag, cidList, auth.Token)
		}
		if err != nil {
			log.Errorf("unbind tag from user err %+v", err)
			return out, err
		}
	}
	out.ResultList = make([]*pb.UnbindTagFromUserResp_Result, 0)
	for cid, success := range resp {
		out.ResultList = append(out.ResultList, &pb.UnbindTagFromUserResp_Result{
			CId:     cid,
			Success: success,
		})
	}
	return out, nil
}

func (s *Server) QueryUserTag(ctx context.Context, in *pb.QueryUserTagReq) (*pb.QueryUserTagResp, error) {
	out := &pb.QueryUserTagResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.mgr.QueryUserTag(uniqueId, in.CId, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			resp, err = s.mgr.QueryUserTag(uniqueId, in.CId, auth.Token)
		}
		if err != nil {
			log.Errorf("query user tag err %+v", err)
			return out, err
		}
	}
	out.TagList = resp
	return out, nil
}

func (s *Server) AddBlackList(ctx context.Context, in *pb.AddBlackListReq) (*pb.AddBlackListResp, error) {
	out := &pb.AddBlackListResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.AddBlackList(uniqueId, in.CIdList, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			err = s.mgr.AddBlackList(uniqueId, in.CIdList, auth.Token)
		}
		if err != nil {
			log.Errorf("add black list err %+v", err)
			return out, err
		}
	}
	return out, nil
}

func (s *Server) DelBlackList(ctx context.Context, in *pb.DelBlackListReq) (*pb.DelBlackListResp, error) {
	out := &pb.DelBlackListResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.DelBlackList(uniqueId, in.CIdList, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			err = s.mgr.DelBlackList(uniqueId, in.CIdList, auth.Token)
		}
		if err != nil {
			log.Errorf("del black list err %+v", err)
			return out, err
		}
	}
	return out, nil
}

func (s *Server) QueryUserStatus(ctx context.Context, in *pb.QueryUserStatusReq) (*pb.QueryUserStatusResp, error) {
	out := &pb.QueryUserStatusResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.mgr.QueryUserStatus(uniqueId, in.CIdList, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			resp, err = s.mgr.QueryUserStatus(uniqueId, in.CIdList, auth.Token)
		}
		if err != nil {
			log.Errorf("query user status err %+v", err)
			return out, err
		}
	}
	out.StatusList = make([]*pb.QueryUserStatusResp_UserStatus, 0)
	for cid, detail := range resp {
		out.StatusList = append(out.StatusList, &pb.QueryUserStatusResp_UserStatus{
			CId:           cid,
			LastLoginTime: detail["last_login_time"],
			Status:        detail["status"],
		})
	}
	return out, nil
}

func (s *Server) QueryDeviceStatus(ctx context.Context, in *pb.QueryDeviceStatusReq) (*pb.QueryDeviceStatusResp, error) {
	out := &pb.QueryDeviceStatusResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.mgr.QueryDeviceStatus(uniqueId, in.CIdList, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			resp, err = s.mgr.QueryDeviceStatus(uniqueId, in.CIdList, auth.Token)
		}
		if err != nil {
			log.Errorf("query device status err %+v", err)
			return out, err
		}
	}
	out.StatusList = make([]*pb.QueryDeviceStatusResp_DeviceStatus, 0)
	for cid, detail := range resp {
		out.StatusList = append(out.StatusList, &pb.QueryDeviceStatusResp_DeviceStatus{
			CId:          cid,
			Availiable:   detail["available"] == "true",
			CIdStatus:    detail["cid_status"],
			DeviceStatus: detail["device_status"],
		})
	}
	return out, nil
}

func (s *Server) QueryUserInfo(ctx context.Context, in *pb.QueryUserInfoReq) (*pb.QueryUserInfoResp, error) {
	out := &pb.QueryUserInfoResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	invalid, validDetail, err := s.mgr.QueryUserInfo(uniqueId, in.CIdList, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			invalid, validDetail, err = s.mgr.QueryUserInfo(uniqueId, in.CIdList, auth.Token)
		}
		if err != nil {
			log.Errorf("query user info err %+v", err)
			return out, err
		}
	}
	out.InvalidList = invalid
	out.UserInfoList = make([]*pb.QueryUserInfoResp_UserInfo, 0)
	for cid, detail := range validDetail {
		phoneType, notificationSwitch, loginFreq := int64(-1), false, int64(-1)
		if val, ok := detail["phone_type"]; ok {
			v, err := strconv.ParseInt(val, 10, 64)
			if err == nil {
				phoneType = v
			}
		}
		if val, ok := detail["notification_switch"]; ok {
			v, err := strconv.ParseBool(val)
			if err == nil {
				notificationSwitch = v
			}
		}
		if val, ok := detail["login_freq"]; ok {
			v, err := strconv.ParseInt(val, 10, 64)
			if err == nil {
				loginFreq = v
			}
		}
		out.UserInfoList = append(out.UserInfoList, &pb.QueryUserInfoResp_UserInfo{
			CId:                cid,
			ClientAppId:        detail["client_app_id"],
			PackageName:        detail["package_name"],
			DeviceToken:        detail["device_token"],
			PhoneType:          int32(phoneType),
			PhoneModel:         detail["phone_model"],
			NotificationSwitch: notificationSwitch,
			CreateTime:         detail["create_time"],
			LoginFreq:          int32(loginFreq),
		})
	}
	return out, nil
}

func (s *Server) SetPushBadge(ctx context.Context, in *pb.SetPushBadgeReq) (*pb.SetPushBadgeResp, error) {
	out := &pb.SetPushBadgeResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CIdList) == 0 || len(in.Operation) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	operation := &Operation{
		Badge: in.Operation,
	}
	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.SetPushBadge(uniqueId, in.CIdList, operation, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			err = s.mgr.SetPushBadge(uniqueId, in.CIdList, operation, auth.Token)
		}
		if err != nil {
			log.Errorf("set push badge err %+v", err)
			return out, err
		}
	}
	return out, nil
}

func (s *Server) QueryUserCount(ctx context.Context, in *pb.QueryUserCountReq) (*pb.QueryUserCountResp, error) {
	out := &pb.QueryUserCountResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TagList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	tagList := &ComplexTagList{
		Tag: make([]*Tag, 0),
	}
	for i := range in.TagList {
		tagList.Tag = append(tagList.Tag, &Tag{
			Key:     in.TagList[i].Key,
			Values:  in.TagList[i].Values,
			OptType: in.TagList[i].OptType,
		})
	}
	auth, err := s.auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	count, err := s.mgr.QueryUserCount(uniqueId, tagList, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err err %+v", err)
				return out, err
			}
			count, err = s.mgr.QueryUserCount(uniqueId, tagList, auth.Token)
		}
		if err != nil {
			log.Errorf("query push count err %+v", err)
			return out, err
		}
	}
	out.Count = int32(count)
	return out, nil
}

func (s *Server) Close() error {
	// nothing
	return nil
}

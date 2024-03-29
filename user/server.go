package user

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	*proto.UnimplementedUserServer
	mgr  *AgentMgr
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
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		log.Errorf("new getui agent err %+v", err)
		return nil, err
	}
	err = mgr.RegisterAgent(core.UniqueId{PushAgent: conf.GeTui.AgentId, BundleId: conf.GeTui.BundleId}, agt)
	if err != nil {
		log.Errorf("register getui agent err %+v", err)
		return nil, err
	}

	auth, err := core.NewAuthCache()
	if err != nil {
		log.Errorf("new auth cache err %+v", err)
		return nil, err
	}

	return &Server{
		mgr:  mgr,
		auth: auth,
	}, nil
}

func (s *Server) BindAlias(ctx context.Context, in *proto.BindAliasReq) (*proto.BindAliasResp, error) {
	out := &proto.BindAliasResp{}

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

	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.BindAlias(ctx, uniqueId, aliasList, auth.Token)
	if err != nil {
		log.Errorf("bind alias err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) QueryAliasByCid(ctx context.Context, in *proto.QueryAliasReq) (*proto.QueryAliasResp, error) {
	out := &proto.QueryAliasResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	alias, err := s.mgr.QueryAliasByCid(ctx, uniqueId, in.CId, auth.Token)
	if err != nil {
		log.Errorf("query alias err %+v", err)
		return out, err
	}
	out.Alias = alias
	return out, nil
}

func (s *Server) QueryCidByAlias(ctx context.Context, in *proto.QueryCidReq) (*proto.QueryCidResp, error) {
	out := &proto.QueryCidResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.Alias) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	cidList, err := s.mgr.QueryCidByAlias(ctx, uniqueId, in.Alias, auth.Token)
	if err != nil {
		log.Errorf("query cid err %+v", err)
		return out, err
	}
	out.CIdList = cidList
	return out, nil
}

func (s *Server) UnbindAlias(ctx context.Context, in *proto.UnbindAliasReq) (*proto.UnbindAliasResp, error) {
	out := &proto.UnbindAliasResp{}

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

	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.UnbindAlias(ctx, uniqueId, aliasList, auth.Token)
	if err != nil {
		log.Errorf("unbind alias err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) RevokeAlias(ctx context.Context, in *proto.RevokeAliasReq) (*proto.RevokeAliasResp, error) {
	out := &proto.RevokeAliasResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.Alias) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.RevokeAlias(ctx, uniqueId, in.Alias, auth.Token)
	if err != nil {
		log.Errorf("revoke alias err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) BindUserWithTag(ctx context.Context, in *proto.BindUserWithTagReq) (*proto.BindUserWithTagResp, error) {
	out := &proto.BindUserWithTagResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CId) == 0 || len(in.TagList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	tagList := &CustomTagList{
		TagList: in.TagList,
	}

	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.BindUserWithTag(ctx, uniqueId, in.CId, tagList, auth.Token)
	if err != nil {
		log.Errorf("bind user with tag err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) BindTagWithUser(ctx context.Context, in *proto.BindTagWithUserReq) (*proto.BindTagWithUserResp, error) {
	out := &proto.BindTagWithUserResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.Tag) == 0 || len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	cidList := &CidList{
		CidList: in.CIdList,
	}

	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.mgr.BindTagWithUser(ctx, uniqueId, in.Tag, cidList, auth.Token)
	if err != nil {
		log.Errorf("bind tag with user err %+v", err)
		return out, err
	}
	out.ResultList = make([]*proto.BindTagWithUserResp_Result, 0)
	for cid, success := range resp {
		out.ResultList = append(out.ResultList, &proto.BindTagWithUserResp_Result{
			CId:     cid,
			Success: success,
		})
	}
	return out, nil
}

func (s *Server) UnbindTagFromUser(ctx context.Context, in *proto.UnbindTagFromUserReq) (*proto.UnbindTagFromUserResp, error) {
	out := &proto.UnbindTagFromUserResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.Tag) == 0 || len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	cidList := &CidList{
		CidList: in.CIdList,
	}

	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.mgr.UnbindTagFromUser(ctx, uniqueId, in.Tag, cidList, auth.Token)
	if err != nil {
		log.Errorf("unbind tag from user err %+v", err)
		return out, err
	}
	out.ResultList = make([]*proto.UnbindTagFromUserResp_Result, 0)
	for cid, success := range resp {
		out.ResultList = append(out.ResultList, &proto.UnbindTagFromUserResp_Result{
			CId:     cid,
			Success: success,
		})
	}
	return out, nil
}

func (s *Server) QueryUserTag(ctx context.Context, in *proto.QueryUserTagReq) (*proto.QueryUserTagResp, error) {
	out := &proto.QueryUserTagResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.mgr.QueryUserTag(ctx, uniqueId, in.CId, auth.Token)
	if err != nil {
		log.Errorf("query user tag err %+v", err)
		return out, err
	}
	out.TagList = resp
	return out, nil
}

func (s *Server) AddBlackList(ctx context.Context, in *proto.AddBlackListReq) (*proto.AddBlackListResp, error) {
	out := &proto.AddBlackListResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.AddBlackList(ctx, uniqueId, in.CIdList, auth.Token)
	if err != nil {
		log.Errorf("add black list err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) DelBlackList(ctx context.Context, in *proto.DelBlackListReq) (*proto.DelBlackListResp, error) {
	out := &proto.DelBlackListResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.DelBlackList(ctx, uniqueId, in.CIdList, auth.Token)
	if err != nil {
		log.Errorf("del black list err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) QueryUserStatus(ctx context.Context, in *proto.QueryUserStatusReq) (*proto.QueryUserStatusResp, error) {
	out := &proto.QueryUserStatusResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.mgr.QueryUserStatus(ctx, uniqueId, in.CIdList, auth.Token)
	if err != nil {
		log.Errorf("query user status err %+v", err)
		return out, err
	}
	out.StatusList = make([]*proto.QueryUserStatusResp_UserStatus, 0)
	for cid, detail := range resp {
		out.StatusList = append(out.StatusList, &proto.QueryUserStatusResp_UserStatus{
			CId:           cid,
			LastLoginTime: detail["last_login_time"],
			Status:        detail["status"],
		})
	}
	return out, nil
}

func (s *Server) QueryDeviceStatus(ctx context.Context, in *proto.QueryDeviceStatusReq) (*proto.QueryDeviceStatusResp, error) {
	out := &proto.QueryDeviceStatusResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.mgr.QueryDeviceStatus(ctx, uniqueId, in.CIdList, auth.Token)
	if err != nil {
		log.Errorf("query device status err %+v", err)
		return out, err
	}
	out.StatusList = make([]*proto.QueryDeviceStatusResp_DeviceStatus, 0)
	for cid, detail := range resp {
		out.StatusList = append(out.StatusList, &proto.QueryDeviceStatusResp_DeviceStatus{
			CId:          cid,
			Available:    detail["available"] == "true",
			CIdStatus:    detail["cid_status"],
			DeviceStatus: detail["device_status"],
		})
	}
	return out, nil
}

func (s *Server) QueryUserInfo(ctx context.Context, in *proto.QueryUserInfoReq) (*proto.QueryUserInfoResp, error) {
	out := &proto.QueryUserInfoResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CIdList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	invalid, validDetail, err := s.mgr.QueryUserInfo(ctx, uniqueId, in.CIdList, auth.Token)
	if err != nil {
		log.Errorf("query user info err %+v", err)
		return out, err
	}
	out.InvalidList = invalid
	out.UserInfoList = make([]*proto.QueryUserInfoResp_UserInfo, 0)
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
		out.UserInfoList = append(out.UserInfoList, &proto.QueryUserInfoResp_UserInfo{
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

func (s *Server) SetPushBadge(ctx context.Context, in *proto.SetPushBadgeReq) (*proto.SetPushBadgeResp, error) {
	out := &proto.SetPushBadgeResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.CIdList) == 0 || len(in.Operation) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	operation := &Operation{
		Badge: in.Operation,
	}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	err = s.mgr.SetPushBadge(ctx, uniqueId, in.CIdList, operation, auth.Token)
	if err != nil {
		log.Errorf("set push badge err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) QueryUserCount(ctx context.Context, in *proto.QueryUserCountReq) (*proto.QueryUserCountResp, error) {
	out := &proto.QueryUserCountResp{}

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
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	count, err := s.mgr.QueryUserCount(ctx, uniqueId, tagList, auth.Token)
	if err != nil {
		log.Errorf("query push count err %+v", err)
		return out, err
	}
	out.Count = int32(count)
	return out, nil
}

func (s *Server) ManageCidAndDeviceToken(ctx context.Context, in *proto.ManageCidAndDeviceTokenReq) (*proto.ManageCidAndDeviceTokenResp, error) {
	out := &proto.ManageCidAndDeviceTokenResp{}

	if len(in.DtList) == 0 || len(in.Manufacturer) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	C2DList := &CidAndDeviceTokenList{
		DTList: make([]*DT, 0, len(in.DtList)),
	}
	for cid, deviceToken := range in.DtList {
		C2DList.DTList = append(C2DList.DTList, &DT{
			Cid:         cid,
			DeviceToken: deviceToken,
		})
	}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	errorList, err := s.mgr.ManageCidAndDeviceToken(ctx, uniqueId, in.Manufacturer, C2DList, auth.Token)
	if err != nil {
		log.Errorf("manage cid and deviceToken err %+v", err)
		return out, err
	}
	for i := range errorList {
		out.ErrorList = append(out.ErrorList, &proto.ManageCidAndDeviceTokenResp_Result{
			Cid:         errorList[i].Cid,
			DeviceToken: errorList[i].DeviceToken,
			ErrorCode:   int32(errorList[i].ErrorCode),
		})
	}
	return out, nil
}

func (s *Server) Close() error {
	s.auth.Close()
	s.mgr.Close()
	return nil
}

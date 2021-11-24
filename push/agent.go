package push

import (
	"fmt"

	"github.com/cestlascorpion/offlinepush/core"
)

type Agent interface {
	PushSingleByCid(request *PushSingleReq, token string) (map[string]map[string]string, error)
	PushSingleByAlias(request *PushSingleReq, token string) (map[string]map[string]string, error)
	PushBatchByCid(request *PushBatchReq, token string) (map[string]map[string]string, error)
	PushBatchByAlias(request *PushBatchReq, token string) (map[string]map[string]string, error)
	CreateMsg(request *CreateMsgReq, token string) (string, error)
	PushListByCid(request *PushListReq, token string) (map[string]map[string]string, error)
	PushListByAlias(request *PushListReq, token string) (map[string]map[string]string, error)
	PushAll(request *PushAllReq, token string) (string, error)
	PushByTag(request *PushByTagReq, token string) (string, error)
	PushByFastCustomTag(request *PushByFastCustomTagReq, token string) (string, error)
	StopPush(taskId, token string) (bool, error)
	DeleteScheduleTask(taskId, token string) (bool, error)
	QueryScheduleTask(taskId, token string) (map[string]string, error)
	QueryDetail(taskId, cId, token string) ([][2]string, error)
}

type AgentMgr struct {
	Agents map[core.UniqueId]Agent
}

func NewAgentMgr() (*AgentMgr, error) {
	return &AgentMgr{
		Agents: make(map[core.UniqueId]Agent),
	}, nil
}

func (m *AgentMgr) RegisterAgent(uniqueId core.UniqueId, agent Agent) error {
	m.Agents[uniqueId] = agent
	return nil
}

func (m *AgentMgr) PushSingle(uniqueId core.UniqueId, request *PushSingleReq, token string) (map[string]map[string]string, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	if len(request.Audience.Cid) != 0 {
		return agent.PushSingleByCid(request, token)
	}
	return agent.PushSingleByAlias(request, token)
}

func (m *AgentMgr) PushBatch(uniqueId core.UniqueId, request *PushBatchReq, token string) (map[string]map[string]string, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	if len(request.MsgList[0].Audience.Cid) != 0 {
		return agent.PushBatchByCid(request, token)
	}
	return agent.PushBatchByAlias(request, token)
}

func (m *AgentMgr) CreateMsg(uniqueId core.UniqueId, request *CreateMsgReq, token string) (string, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return "", fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.CreateMsg(request, token)
}

func (m *AgentMgr) PushList(uniqueId core.UniqueId, request *PushListReq, token string) (map[string]map[string]string, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	if len(request.Audience.Cid) != 0 {
		return agent.PushListByCid(request, token)
	}
	return agent.PushListByAlias(request, token)
}

func (m *AgentMgr) PushAll(uniqueId core.UniqueId, request *PushAllReq, token string) (string, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return "", fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.PushAll(request, token)
}

func (m *AgentMgr) PushByTag(uniqueId core.UniqueId, request *PushByTagReq, token string) (string, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return "", fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.PushByTag(request, token)
}

func (m *AgentMgr) PushByFastCustomTag(uniqueId core.UniqueId, request *PushByFastCustomTagReq, token string) (string, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return "", fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.PushByFastCustomTag(request, token)
}

func (m *AgentMgr) StopPush(uniqueId core.UniqueId, taskId, token string) (bool, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return false, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.StopPush(taskId, token)
}

func (m *AgentMgr) DeleteScheduleTask(uniqueId core.UniqueId, taskId, token string) (bool, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return false, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.DeleteScheduleTask(taskId, token)
}

func (m *AgentMgr) QueryScheduleTask(uniqueId core.UniqueId, taskId, token string) (map[string]string, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryScheduleTask(taskId, token)
}

func (m *AgentMgr) QueryDetail(uniqueId core.UniqueId, taskId, cId, token string) ([][2]string, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryDetail(taskId, cId, token)
}

type Audience struct {
	Cid           []string `json:"cid,omitempty"`             // cid数组，单推只能填一个cid，批量推可以填写多个（数组长度小于200）
	Alias         []string `json:"alias,omitempty"`           // 别名数组，单推只能填一个别名，批量推可以填写多个（数组长度小于200）
	Tag           *[]Tag   `json:"tag,omitempty"`             // 推送条件
	FastCustomTag string   `json:"fast_custom_tag,omitempty"` // 使用用户标签筛选目标用户
}

type Tag struct {
	Key     string   `json:"key"`      // 必须字段，默认值：无，查询条件 phone_type 手机类型; region 省市; custom_tag 用户标签; portrait 个推用户画像使用编码
	Values  []string `json:"values"`   // 必须字段，默认值：无，查询条件值列表，其中 手机型号使用android和ios； 省市使用编号，
	OptType string   `json:"opt_type"` // 必须字段，默认值：无，or(或)，and(与)，not(非)，values间的交并补操作
	/*
	 * 不同key之间是交集，同一个key之间是根据opt_type操作
	 * 需要发送给城市在A/B/C里面，没有设置tagtest标签，手机型号为android的用户，用条件交并补功能可以实现，city(A|B|C) && !tag(tagtest) && phonetype(android)
	 */
}

type Settings struct {
	TTL          int64     `json:"ttl,omitempty"`           // 非必须，默认一小时，消息离线时间设置，单位毫秒，-1表示不设离线，-1 ～ 3 * 24 * 3600 * 1000(3天)之间
	Strategy     *Strategy `json:"strategy,omitempty"`      // 非必须，默认值：{"strategy":{"default":1}}，厂商通道策略
	Speed        int       `json:"speed,omitempty"`         // 非必须，定速推送，例如100，个推控制下发速度在100条/秒左右，0表示不限速
	ScheduleTime int64     `json:"schedule_time,omitempty"` // 非必须，定时推送时间，必须是7天内的时间，格式：毫秒时间戳
}

type Strategy struct {
	Default int `json:"default,omitempty"`
	/*
	 * default字段，非必须，默认值为 1
	 * 默认所有通道的策略选择1-4
	 * 1: 表示该消息在用户在线时推送个推通道，用户离线时推送厂商通道;
	 * 2: 表示该消息只通过厂商通道策略下发，不考虑用户是否在线;
	 * 3: 表示该消息只通过个推通道下发，不考虑用户是否在线；
	 * 4: 表示该消息优先从厂商通道下发，若消息内容在厂商通道代发失败后会从个推通道下发。
	 * 其中名称可填写: ios、st、hw、xm、vv、mz、op，
	 */
	Ios int `json:"ios,omitempty"` // 非必须，ios通道策略1-4，表示含义同上，要推送ios通道，需要在个推开发者中心上传ios证书，建议填写2或4，否则可能会有消息不展示的问题
	St  int `json:"st,omitempty"`  // 非必须，通道策略1-4，表示含义同上，需要开通st厂商使用该通道推送消息
	Hw  int `json:"hw,omitempty"`  // 非必须，通道策略1-4，表示含义同上
	Xm  int `json:"xm,omitempty"`  // 非必须，通道策略1-4，表示含义同上
	Vv  int `json:"vv,omitempty"`  // 非必须，通道策略1-4，表示含义同上
	Mz  int `json:"mz,omitempty"`  // 非必须，通道策略1-4，表示含义同上
	Op  int `json:"op,omitempty"`  // 非必须，通道策略1-4，表示含义同上
}

type PushMessage struct {
	Duration string `json:"duration,omitempty"`
	/*
	 * duration字段，非必须
	 * 手机端通知展示时间段，格式为毫秒时间戳段，两个时间的时间差必须大于10分钟
	 * 例如："1590547347000-1590633747000"
	 */
	Notification *PushMessageNotification `json:"notification,omitempty"` // 非必须，通知消息内容，仅支持安卓系统，iOS系统不展示个推通知消息，与transmission、revoke三选一，都填写时报错
	Transmission string                   `json:"transmission,omitempty"` // 非必须，纯透传消息内容，安卓和iOS均支持，与notification、revoke 三选一，都填写时报错，长度 ≤ 3072
	Revoke       *PushMessageRevoke       `json:"revoke,omitempty"`       // 非必须，撤回消息时使用，与notification、transmission三选一，都填写时报错
}

type PushMessageNotification struct {
	Title        string `json:"title"`                  // 必须，通知消息标题，长度 ≤ 50
	Body         string `json:"body"`                   // 必须，通知消息内容，长度 ≤ 256
	BigText      string `json:"big_text,omitempty"`     // 非必须，长文本消息内容，通知消息+长文本样式，与big_image二选一，两个都填写时报错，长度 ≤ 512
	BigImage     string `json:"big_image,omitempty"`    // 非必须，大图的URL地址，通知消息+大图样式， 与big_text二选一，两个都填写时报错，长度 ≤ 1024
	Logo         string `json:"logo,omitempty"`         // 非必须，通知的图标名称，包含后缀名（需要在客户端开发时嵌入），如“push.png”，长度 ≤ 64
	LogoUrl      string `json:"logo_url,omitempty"`     // 非必须，通知图标URL地址，长度 ≤ 256
	ChannelId    string `json:"channel_id,omitempty"`   // 非必须，默认值：Default，通知渠道id，长度 ≤ 64
	ChannelName  string `json:"channel_name,omitempty"` // 非必须，默认值：Default，通知渠道名称，长度 ≤ 64
	ChannelLevel int    `json:"channel_level,omitempty"`
	/*
	 * channel_level，非必须，默认值：3
	 * 设置通知渠道重要性（可以控制响铃，震动，浮动，闪灯等等）
	 * android8.0以下
	 * 0，1，2:无声音，无振动，不浮动
	 * 3:有声音，无振动，不浮动
	 * 4:有声音，有振动，有浮动
	 * android8.0以上
	 * 0：无声音，无振动，不显示；
	 * 1：无声音，无振动，锁屏不显示，通知栏中被折叠显示，导航栏无logo;
	 * 2：无声音，无振动，锁屏和通知栏中都显示，通知不唤醒屏幕;
	 * 3：有声音，无振动，锁屏和通知栏中都显示，通知唤醒屏幕;
	 * 4：有声音，有振动，亮屏下通知悬浮展示，锁屏通知以默认形式展示且唤醒屏幕;
	 */
	ClickType string `json:"click_type,omitempty"`
	/*
	 * click_type，必须，默认值：无
	 * 点击通知后续动作，目前支持以下后续动作：
	 * intent：打开应用内特定页面
	 * url：打开网页地址
	 * payload：自定义消息内容启动应用
	 * payload_custom：自定义消息内容不启动应用
	 * startapp：打开应用首页
	 * none：纯通知，无后续动作
	 */
	Intent string `json:"intent,omitempty"`
	/*
	 * click_type为intent时必填
	 * 点击通知打开应用特定页面，长度 ≤ 2048
	 * 示例：intent:#Intent;component=你的包名/你要打开的 activity 全路径;S.parm1=value1;S.parm2=value2;end
	 */
	Url         string `json:"url,omitempty"`       // click_type为url时必填，点击通知打开链接，长度 ≤ 1024
	PayLoad     string `json:"payload,omitempty"`   // click_type为payload/payload_custom时必填，点击通知加自定义消息，长度 ≤ 3072
	NotifyId    int64  `json:"notify_id,omitempty"` // 非必须，覆盖任务时会使用到该字段，两条消息的notify_id相同，新的消息会覆盖老的消息，范围：0-2147483647
	RingName    string `json:"ring_name,omitempty"` // 非必须，自定义铃声，请填写文件名，不包含后缀名(需要在客户端开发时嵌入)，个推通道下发有效 客户端SDK最低要求 2.14.0.0
	BadgeAddNum int    `json:"badge_add_num,omitempty"`
	/*
	 * badge_add_num，非必须，默认值：无
	 * 角标，必须大于0，个推通道下发有效
	 * 此属性目前仅针对华为 EMUI 4.1 及以上设备有效
	 * 角标数字数据会和之前角标数字进行叠加，也就是相加；
	 * 举例：角标数字配置1，应用之前角标数为2，发送此角标消息后，应用角标数显示为3。
	 * 客户端SDK最低要求 2.14.0.0
	 */
	ThreadId string `json:"thread_id,omitempty"` // 非必须，消息折叠分组，设置成相同thread_id的消息会被折叠（仅支持个推渠道下发的安卓消息）。目前与iOS的thread_id设置无关，安卓和iOS需要分别设置。
}

type PushMessageRevoke struct {
	OldTaskId string `json:"old_task_id"`     // 必须，需要撤回的taskId
	Force     bool   `json:"force,omitempty"` // 非必须，【小心使用】在没有找到对应的taskId，是否把对应appId下所有的通知都撤回
}

type PushChannel struct {
	Ios     *IosChannel     `json:"ios,omitempty"`     // 非必须，ios通道推送消息内容
	Android *AndroidChannel `json:"android,omitempty"` // 非必须，android通道推送消息内容
}

type IosChannel struct {
	// 具体参数含义详见苹果APNs文档
	// https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/PayloadKeyReference.html
	Type           string        `json:"type,omitempty"`             // 非必须，默认值：notify，voip：voip语音推送，notify：apns通知消息
	Aps            *Aps          `json:"aps,omitempty"`              // 推送通知消息内容
	AutoBadge      string        `json:"auto_badge,omitempty"`       // 非必须，用于计算icon上显示的数字，还可以实现显示数字的自动增减，如“+1”、 “-1”、 “1” 等，计算结果将覆盖badge
	PayLoad        string        `json:"payload,omitempty"`          // 非必须，增加自定义的数据
	Multimedia     *[]Multimedia `json:"multimedia,omitempty"`       // 非必须，该字段为Array类型，设置多媒体
	ApnsCollapseId string        `json:"apns-collapse-id,omitempty"` // 非必须，使用相同的apns-collapse-id可以覆盖之前的消息
}

type Aps struct {
	Alert            *Alert `json:"alert,omitempty"` // 非必须，通知消息
	ContentAvailable int    `json:"content-available,omitempty"`
	/*
	 * content-available非必须，默认值：0
	 * 0表示普通通知消息(默认为0)
	 * 1表示静默推送(无通知栏消息)，静默推送时不需要填写其他参数
	 * 苹果建议1小时最多推送3条静默消息
	 */
	Sound    string `json:"sound,omitempty"`     // 非必须，通知铃声文件名，如果铃声文件未找到，响铃为系统默认铃声。 无声设置为“com.gexin.ios.silence”或不填
	Category string `json:"category,omitempty"`  // 非必须，在客户端通知栏触发特定的action和button显示
	ThreadId string `json:"thread-id,omitempty"` // 非必须，ios的远程通知通过该属性对通知进行分组，仅支持iOS 12.0以上版本
}

type Alert struct {
	Title           string   `json:"title,omitempty"`             // 非必须，通知消息标题
	Body            string   `json:"body,omitempty"`              // 非必须，通知消息内容
	ActionLocKey    string   `json:"action-loc-key,omitempty"`    // 非必须，（用于多语言支持）指定执行按钮所使用的Localizable.strings
	LocKey          string   `json:"loc-key,omitempty"`           // 非必须，（用于多语言支持）指定Localizable.strings文件中相应的key
	LocArgs         []string `json:"loc-args,omitempty"`          // 非必须，如果loc-key中使用了占位符，则在loc-args中指定各参数
	LaunchImage     string   `json:"launch-image,omitempty"`      // 非必须，指定启动界面图片名
	TitleLocKey     string   `json:"title-loc-key,omitempty"`     // 非必须，(用于多语言支持）对于标题指定执行按钮所使用的Localizable.strings，仅支持iOS8.2以上版本
	TitleLocArgs    []string `json:"title-loc-args,omitempty"`    // 非必须，对于标题，如果loc-key中使用的占位符，则在loc-args中指定各参数，仅支持iOS8.2以上版本
	SubTitle        string   `json:"sub_title,omitempty"`         // 非必须，通知子标题，仅支持iOS8.2以上版本
	SubTitleLocKey  string   `json:"subtitle-loc-key,omitempty"`  // 非必须，当前本地化文件中的子标题字符串的关键字，仅支持iOS8.2以上版本
	SubTitleLocArgs []string `json:"subtitle-loc-args,omitempty"` // 非必须，当前本地化子标题内容中需要置换的变量参数，仅支持iOS8.2以上版本
}

type Multimedia struct {
	Url      string `json:"url"`                 // 必须，多媒体资源地址
	Type     int    `json:"type"`                // 必须，资源类型（1.图片，2.音频，3.视频）
	OnlyWifi bool   `json:"only_wifi,omitempty"` // 非必须，是否只在wifi环境下加载，如果设置成true，但未使用wifi时，会展示成普通通知
}

type AndroidChannel struct {
	Ups *Ups `json:"ups"` // android厂商通道推送消息内容
}

type Ups struct {
	Notification *PushChannelNotification `json:"notification,omitempty"` // 非必须，通知消息内容，与transmission、revoke三选一，都填写时报错
	TransMission string                   `json:"transmission,omitempty"` // 非必须，通知消息内容，与transmission、revoke三选一，都填写时报错
	Revoke       *PushChannelRevoke       `json:"revoke,omitempty"`       // 非必须，撤回消息时使用(仅撤回厂商通道消息，支持的厂商有小米、VIVO)，与notification、transmission三选一，都填写时报错(消息撤回请勿填写策略参数)
	Options      *Options                 `json:"options,omitempty"`      // 非必须，三方厂商扩展内容
}

type PushChannelNotification struct {
	Title string `json:"title"`
	/*
	 * 必须，通知栏标题（长度建议取最小集）
	 * 小米：title长度限制为50字
	 * 华为：title长度限制40字
	 * 魅族：title长度限制32字
	 * OPPO：title长度限制32字
	 * VIVO：title长度限制40英文字符
	 */
	Body string `json:"body"`
	/*
	 * 必须，通知栏内容(长度建议取最小集)
	 * 小米：content长度限制128字
	 * 华为：content长度小于1024字
	 * 魅族：content长度限制100字
	 * OPPO：content长度限制200字
	 * VIVO：content长度限制100个英文字符
	 */
	ClickType string `json:"click_type,omitempty"`
	/*
	 * click_type，必须，点击通知后续动作，目前支持以下后续动作，
	 * intent：打开应用内特定页面(厂商都支持)，
	 * url：打开网页地址(厂商都支持，华为要求https协议)，
	 * startapp：打开应用首页(厂商都支持)
	 */
	Intent string `json:"intent,omitempty"`
	/*
	 * click_type为intent时必填
	 * 点击通知打开应用特定页面，intent格式必须正确且不能为空，长度 ≤ 4096;【注意：vivo侧厂商限制 ≤ 1024】
	 * 示例：intent:#Intent;component=你的包名/你要打开的 activity 全路径;S.parm1=value1;S.parm2=value2;end
	 */
	Url      string `json:"url,omitempty"`       // click_type为url时必填，点击通知打开链接，长度 ≤ 1024
	PayLoad  string `json:"payload,omitempty"`   // click_type为payload/payload_custom时必填，点击通知加自定义消息，长度 ≤ 3072
	NotifyId int64  `json:"notify_id,omitempty"` // 非必须，消息覆盖使用，两条消息的notify_id相同，新的消息会覆盖老的消息，范围：0-2147483647
}

type PushChannelRevoke struct {
	OldTaskId string `json:"old_task_id"` // 必须，需要撤回的taskId
}

type Options struct {
	HwOptions *HwOptions `json:"HW,omitempty"`
	XmOptions *XmOptions `json:"XM,omitempty"`
	VvOptions *VvOptions `json:"VV,omitempty"`
	OpOptions *OpOptions `json:"OP,omitempty"`
}

type HwOptions struct {
	Class        string `json:"/message/android/notification/badge/class,omitempty"`   // value：应用入口Activity路径名称
	AddNum       int    `json:"/message/android/notification/badge/add_num,omitempty"` // value：应用角标累加数字，并非应用角标实际显示数字，必须是大于0小于100的整数
	SetNum       int    `json:"/message/android/notification/badge/set_num,omitempty"` // value：角标设置数字，必须是大于0小于100的整数，如果set_num和add_num同时存在，以set_num为准
	Image        string `json:"/message/android/notification/image,omitempty"`         // 通知小图；value：请写入对应图标https地址，URL使用的协议必须是HTTPS协议
	Style        int    `json:"/message/android/notification/style,omitempty"`         // 通知长文本；value：1，通知栏样式 1，为长文本。
	BigTitle     string `json:"/message/android/notification/big_title,omitempty"`     // 通知长文本；value：通知title内容，设置big_title后通知栏展示时，使用 big_title而不用title。
	BigBody      string `json:"/message/android/notification/big_body,omitempty"`      // 通知长文本；value：通知title内容，通知body内容。设置big_body后通知栏展示时，使用big_body而不用body。
	Importance   string `json:"/message/android/notification/importance,omitempty"`    // 取值为“LOW”时，表示消息为资讯营销；取值为“NORMAL”时，表示消息为服务与通讯
	DefaultSound bool   `json:"/message/android/notification/default_sound,omitempty"` // 设置为 false，使用sound自定义铃声
	ChannelId    string `json:"/message/android/notification/channel_id,omitempty"`    // 自Android O版本后可以支持通知栏自定义渠道，指定消息要展示在哪个通知渠道上，详情请参见自定义通知渠道。 自定义通知渠道仅对发送给用户设备的重要级别消息有效，一般级别消息仍然通过华为营销通知渠道展示。
	Sound        string `json:"/message/android/notification/sound,omitempty"`         // 自定义消息通知铃声，在新创建渠道时有效，此处设置的铃声文件必须存放在应用的/res/raw路径下，例如设置为“/raw/shake”，对应应用本地的“/res/raw/shake.xxx”文件，支持的文件格式包括MP3、WAV、MPEG等，如果不设置使用默认系统铃声。
}

type XmOptions struct {
	Channel      string `json:"channel,omitempty"`                     // value：填写小米平台申请的渠道id
	LargeIconUri string `json:"notification_large_icon_uri,omitempty"` // 通知小图； value：填写接口返回的图片链接。Large icon可以出现在大图版和多字版消息中，显示在右边，通知小图。图片要求:尺寸必须为 120×120px，文件小于200KB，PNG/JPG/JPEG格式。
	StyleType    string `json:"notification_style_type,omitempty"`     // 通知大文本； value：1，大文本和大图二选一。
	BigTxt       string `json:"notification_bigTxt,omitempty"`         // 通知大文本； value：文本内容，支持展示多行文字；最多128字，可以通过\n换行。
	BigPic       string `json:"notification_bigPic_uri,omitempty"`     // 通知大文本； value：填写接口返回的图片链接，图片要求:固定876x324px，小于1M，PNG/JPG/JPEG格式。
	SoundUri     string `json:"sound_uri,omitempty"`                   // 小米后台申请的自定义 sound_url 地址，示例：android.resource://your packagename/raw/XXX
}

type VvOptions struct {
	Classification int `json:"classification,omitempty"` // value: 0 代表运营消息，1代表系统消息，不填默认为0
}

type OpOptions struct {
	Channel string `json:"channel,omitempty"` // value：填写OPPO平台登记的渠道ID
}

type PushSingleReq struct {
	RequestId   string       `json:"request_id"`             // 必须字段，请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	Audience    *Audience    `json:"audience"`               // 必须字段，cid数组，只能填一个cid
	Settings    *Settings    `json:"settings,omitempty"`     // 非必须，推送条件设置
	PushMessage *PushMessage `json:"push_message"`           // 必须字段，个推推送消息参数
	PushChannel *PushChannel `json:"push_channel,omitempty"` // 非必须，厂商推送消息参数，包含ios消息参数，android厂商消息参数
}

type PushBatchReq struct {
	IsAsync bool             `json:"is_async,omitempty"` // 非必须，默认值：false，是否异步推送，异步推送不会返回data，is_async为false时返回data
	MsgList []*PushSingleReq `json:"msg_list"`           // 必须，默认值：无，消息内容，数组长度不大于 200
}

type CreateMsgReq struct {
	RequestId   string       `json:"request_id,omitempty"`   // 非必须，请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	GroupName   string       `json:"group_name,omitempty"`   // 非必须，任务组名
	Settings    *Settings    `json:"settings,omitempty"`     // 非必须，推送条件设置
	PushMessage *PushMessage `json:"push_message"`           // 必须字段，个推推送消息参数
	PushChannel *PushChannel `json:"push_channel,omitempty"` // 非必须，厂商推送消息参数，包含ios消息参数，android厂商消息参数
}

type PushListReq struct {
	Audience *Audience `json:"audience"`           // 必须字段，用cid数组，多个cid，注意这里！！数组长度不大于200
	IsAsync  bool      `json:"is_async,omitempty"` // 非必须，默认值：false，是否异步推送，异步推送不会返回data，is_async为false时返回data
	TaskId   string    `json:"taskid"`             // 必须字段，默认值：无，使用创建消息接口返回的taskId，可以多次使用
}

type PushAllReq struct {
	RequestId   string       `json:"request_id"`             // 必须，请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	GroupName   string       `json:"group_name,omitempty"`   // 非必须，任务组名
	Audience    string       `json:"audience"`               // 必须字段，必须为all
	Settings    *Settings    `json:"settings,omitempty"`     // 非必须，推送条件设置
	PushMessage *PushMessage `json:"push_message"`           // 必须字段，个推推送消息参数
	PushChannel *PushChannel `json:"push_channel,omitempty"` // 非必须，厂商推送消息参数，包含ios消息参数，android厂商消息参数
}

type PushByTagReq struct {
	RequestId   string       `json:"request_id"`             // 必须，请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	GroupName   string       `json:"group_name,omitempty"`   // 非必须，任务组名
	Settings    *Settings    `json:"settings,omitempty"`     // 非必须，推送条件设置
	Audience    *Audience    `json:"audience"`               // 必须字段，tag数组
	PushMessage *PushMessage `json:"push_message"`           // 必须字段，个推推送消息参数
	PushChannel *PushChannel `json:"push_channel,omitempty"` // 非必须，厂商推送消息参数，包含ios消息参数，android厂商消息参数
}

type PushByFastCustomTagReq struct {
	RequestId   string       `json:"request_id"`             // 必须字段，请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	Settings    *Settings    `json:"settings,omitempty"`     // 非必须，推送条件设置
	Audience    *Audience    `json:"audience"`               // 必须字段，tag数组
	PushMessage *PushMessage `json:"push_message"`           // 必须字段，个推推送消息参数
	PushChannel *PushChannel `json:"push_channel,omitempty"` // 非必须，厂商推送消息参数，包含ios消息参数，android厂商消息参数
}

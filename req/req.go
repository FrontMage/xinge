package req

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/FrontMage/xinge"
)

var XingeURL = "https://openapi.xg.qq.com/v3/push/app"

// URL 修改信鸽的请求URL
func URL(url string) {
	XingeURL = url
}

// ReqOpt 请求选项修改函数
type ReqOpt func(*xinge.Request)

// NewSingleIOSAccountPush 新建一个iOS单账号push请求
func NewSingleIOSAccountPush(
	account, title, content string,
	opts ...ReqOpt,
) (*http.Request, error) {
	req := &xinge.Request{
		Platform:     xinge.PlatformiOS,
		MessageType:  xinge.MsgTypeNotify,
		AudienceType: xinge.AdAccount,
		AccountList:  []string{account},
		Message: xinge.Message{
			Title:   title,
			Content: content,
			IOS: &xinge.IOSParams{
				Aps: &xinge.Aps{
					Alert: map[string]string{
						"title":   title,
						"content": content,
					},
					Badge: 1,
					Sound: "default",
				},
			},
		},
	}
	return NewPushReq(req, opts...)
}

// NewSinglAndroidAccountPush 新建一个安卓通知栏push请求
func NewSinglAndroidAccountPush(
	account, title, content string,
	opts ...ReqOpt,
) (*http.Request, error) {
	req := &xinge.Request{
		Platform:     xinge.PlatformAndroid,
		MessageType:  xinge.MsgTypeNotify,
		AudienceType: xinge.AdAccount,
		AccountList:  []string{account},
		Message: xinge.Message{
			Title:   title,
			Content: content,
		},
	}
	return NewPushReq(req, opts...)
}

// NewPushReq 新建一个push请求
func NewPushReq(req *xinge.Request, opts ...ReqOpt) (request *http.Request, err error) {
	for _, opt := range opts {
		opt(req)
	}
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	request, err = http.NewRequest("POST", XingeURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	return
}

// PlatformIOS 修改push平台为iOS
func PlatformIOS() ReqOpt {
	return func(r *xinge.Request) {
		r.Platform = xinge.PlatformiOS
	}
}

// PlatformAndroid 修改push平台为PlatformAndroid
func PlatformAndroid() ReqOpt {
	return func(r *xinge.Request) {
		r.Platform = xinge.PlatformAndroid
	}
}

// PlatformAll 修改push平台为全量推送
func PlatformAll() ReqOpt {
	return func(r *xinge.Request) {
		r.Platform = xinge.PlatformAll
	}
}

// EnvProd 修改请求环境为product，只对iOS有效
func EnvProd() ReqOpt {
	return func(r *xinge.Request) {
		r.Environment = xinge.EnvProd
	}
}

// EnvDev 修改请求环境为dev，只对iOS有效
func EnvDev() ReqOpt {
	return func(r *xinge.Request) {
		r.Environment = xinge.EnvDev
	}
}

// Title 修改push title
func Title(t string) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Title = t
		if r.Message.IOS != nil {
			if r.Message.IOS.Aps != nil {
				r.Message.IOS.Aps.Alert["title"] = t
			} else {
				r.Message.IOS.Aps = &xinge.Aps{
					Alert: map[string]string{"title": t},
				}
			}
		} else {
			r.Message.IOS = &xinge.IOSParams{
				Aps: &xinge.Aps{
					Alert: map[string]string{"title": t},
				},
			}
		}
	}
}

// Content 修改push content
func Content(c string) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Content = c
		if r.Message.IOS != nil {
			if r.Message.IOS.Aps != nil {
				r.Message.IOS.Aps.Alert["body"] = c
			} else {
				r.Message.IOS.Aps = &xinge.Aps{
					Alert: map[string]string{"body": c},
				}
			}
		} else {
			r.Message.IOS = &xinge.IOSParams{
				Aps: &xinge.Aps{
					Alert: map[string]string{"body": c},
				},
			}
		}
	}
}

// TODO: accept_time modify

// NID 修改nid
func NID(id int) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Android.NID = id
	}
}

// BuilderID 修改builder_id
func BuilderID(id int) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Android.BuilderID = id
	}
}

// Ring 修改ring
func Ring(ring int) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Android.Ring = ring
	}
}

// RingRaw 修改ring_raw
func RingRaw(rr string) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Android.RingRaw = rr
	}
}

// Vibrate 修改vibrate
func Vibrate(v int) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Android.Vibrate = v
	}
}

// Lights 修改lights
func Lights(l int) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Android.Lights = l
	}
}

// Clearable 修改clearable
func Clearable(c int) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Android.Clearable = c
	}
}

// IconType 修改icon_type
func IconType(it int) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Android.IconType = it
	}
}

// IconRes 修改icon_res
func IconRes(ir string) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Android.IconRes = ir
	}
}

// StyleID 修改style_id
func StyleID(s int) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Android.StyleID = s
	}
}

// SmallIcon 修改small_icon
func SmallIcon(si int) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Android.SmallIcon = si
	}
}

// Action 修改action
func Action(a map[string]interface{}) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Android.Action = a
	}
}

// TODO: append action

// CustomContent 修改custom_content 和 custom
func CustomContent(ct map[string]string) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.Android.CustomContent = ct
		r.Message.IOS.Custom = ct
	}
}

// CustomContentSet 设置custom_content和custom的某个字段
func CustomContentSet(k, v string) ReqOpt {
	return func(r *xinge.Request) {
		if r.Message.Android.CustomContent == nil {
			r.Message.Android.CustomContent = map[string]string{k: v}
		} else {
			r.Message.Android.CustomContent[k] = v
		}
		if r.Message.IOS.Custom == nil {
			r.Message.IOS.Custom = map[string]string{k: v}
		} else {
			r.Message.IOS.Custom[k] = v
		}
	}
}

// Aps 修改aps
func Aps(aps *xinge.Aps) ReqOpt {
	return func(r *xinge.Request) {
		r.Message.IOS.Aps = aps
	}
}

// AudienceType 修改audience_type
func AudienceType(at xinge.AudienceType) ReqOpt {
	return func(r *xinge.Request) {
		r.AudienceType = at
	}
}

// Platform 修改platform
func Platform(p xinge.Platform) ReqOpt {
	return func(r *xinge.Request) {
		r.Platform = p
	}
}

// Message 修改message
func Message(m xinge.Message) ReqOpt {
	return func(r *xinge.Request) {
		r.Message = m
	}
}

// TagList 修改tag_list
func TagList(tl *xinge.TagList) ReqOpt {
	return func(r *xinge.Request) {
		r.TagList = tl
	}
}

// TokenList 修改token_list
func TokenList(tl []string) ReqOpt {
	return func(r *xinge.Request) {
		r.TokenList = tl
	}
}

// TokenListAdd 给token_list添加一个token
func TokenListAdd(t string) ReqOpt {
	return func(r *xinge.Request) {
		if r.TokenList != nil {
			r.TokenList = append(r.TokenList, t)
		} else {
			r.TokenList = []string{t}
		}
	}
}

// AccountList 修改account_list
func AccountList(al []string) ReqOpt {
	return func(r *xinge.Request) {
		r.AccountList = al
	}
}

// AccountListAdd 给account_list添加一个account
func AccountListAdd(a string) ReqOpt {
	return func(r *xinge.Request) {
		if r.AccountList != nil {
			r.AccountList = append(r.AccountList, a)
		} else {
			r.AccountList = []string{a}
		}
	}
}

// ExpireTime 修改expire_time
func ExpireTime(et time.Time) ReqOpt {
	return func(r *xinge.Request) {
		r.ExpireTime = int(et.Unix())
	}
}

// SendTime 修改send_time
func SendTime(st time.Time) ReqOpt {
	return func(r *xinge.Request) {
		// TODO: format this time as yyyy-MM-DD HH:MM:SS
		r.SendTime = st.Format("")
	}
}

// MultiPkg 修改multi_pkg
func MultiPkg(mp bool) ReqOpt {
	return func(r *xinge.Request) {
		r.MultiPkg = mp
	}
}

// LoopTimes 修改loop_times
func LoopTimes(lt int) ReqOpt {
	return func(r *xinge.Request) {
		r.LoopTimes = lt
	}
}

// StatTag 修改stat_tag
func StatTag(st string) ReqOpt {
	return func(r *xinge.Request) {
		r.StatTag = st
	}
}

// Seq 修改seq
func Seq(s int64) ReqOpt {
	return func(r *xinge.Request) {
		r.Seq = s
	}
}

// AudienceType 修改account_type
func AccountType(at int) ReqOpt {
	return func(r *xinge.Request) {
		r.AccountType = at
	}
}

// PushID 修改push_id
func PushID(pid string) ReqOpt {
	return func(r *xinge.Request) {
		r.PushID = pid
	}
}

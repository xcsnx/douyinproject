// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoserver

import (
	"context"
	douyinvideo "douyin-user/idl/douyin_video/kitex_gen/douyinvideo"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServerServiceInfo
}

var videoServerServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoServer"
	handlerType := (*douyinvideo.VideoServer)(nil)
	methods := map[string]kitex.MethodInfo{
		"Publish": kitex.NewMethodInfo(publishHandler, newVideoServerPublishArgs, newVideoServerPublishResult, false),
		"GetList": kitex.NewMethodInfo(getListHandler, newVideoServerGetListArgs, newVideoServerGetListResult, false),
		"Feed":    kitex.NewMethodInfo(feedHandler, newVideoServerFeedArgs, newVideoServerFeedResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "douyinvideo",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func publishHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinvideo.VideoServerPublishArgs)
	realResult := result.(*douyinvideo.VideoServerPublishResult)
	success, err := handler.(douyinvideo.VideoServer).Publish(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServerPublishArgs() interface{} {
	return douyinvideo.NewVideoServerPublishArgs()
}

func newVideoServerPublishResult() interface{} {
	return douyinvideo.NewVideoServerPublishResult()
}

func getListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinvideo.VideoServerGetListArgs)
	realResult := result.(*douyinvideo.VideoServerGetListResult)
	success, err := handler.(douyinvideo.VideoServer).GetList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServerGetListArgs() interface{} {
	return douyinvideo.NewVideoServerGetListArgs()
}

func newVideoServerGetListResult() interface{} {
	return douyinvideo.NewVideoServerGetListResult()
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinvideo.VideoServerFeedArgs)
	realResult := result.(*douyinvideo.VideoServerFeedResult)
	success, err := handler.(douyinvideo.VideoServer).Feed(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServerFeedArgs() interface{} {
	return douyinvideo.NewVideoServerFeedArgs()
}

func newVideoServerFeedResult() interface{} {
	return douyinvideo.NewVideoServerFeedResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Publish(ctx context.Context, req *douyinvideo.PublishRequest) (r *douyinvideo.PublishResponse, err error) {
	var _args douyinvideo.VideoServerPublishArgs
	_args.Req = req
	var _result douyinvideo.VideoServerPublishResult
	if err = p.c.Call(ctx, "Publish", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetList(ctx context.Context, req *douyinvideo.GetListRequest) (r *douyinvideo.GetListResponse, err error) {
	var _args douyinvideo.VideoServerGetListArgs
	_args.Req = req
	var _result douyinvideo.VideoServerGetListResult
	if err = p.c.Call(ctx, "GetList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Feed(ctx context.Context, req *douyinvideo.FeedRequest) (r *douyinvideo.FeedResponse, err error) {
	var _args douyinvideo.VideoServerFeedArgs
	_args.Req = req
	var _result douyinvideo.VideoServerFeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

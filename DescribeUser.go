package main

//import (
//	"github.com/yunify/qingcloud-sdk-go/request/data"
//	"github.com/yunify/qingcloud-sdk-go/request"
//)
//
//// Documentation URL: https://docs.qingcloud.com/api/image/describe-image-users.html
//func (s *Proxy) DescribeImageUsers(i *DescribeImageUsersInput) (*DescribeImageUsersOutput, error) {
//	if i == nil {
//		i = &DescribeImageUsersInput{}
//	}
//	o := &data.Operation{
//		Config:        &s.Config,
//		Properties:    s.Properties,
//		APIName:       "DescribeImageUsers",
//		RequestMethod: "GET",
//	}
//
//	x := &DescribeImageUsersOutput{}
//	r, err := request.New(o, i, x)
//	if err != nil {
//		return nil, err
//	}
//
//	err = r.Send()
//	if err != nil {
//		return nil, err
//	}
//
//	return x, err
//}
//
//type DescribeImageUsersInput struct {
//	ImageID *string `json:"image_id" name:"image_id" location:"params"` // Required
//	Limit   *int    `json:"limit" name:"limit" default:"20" location:"params"`
//	Offset  *int    `json:"offset" name:"offset" default:"0" location:"params"`
//}
//
//func (v *DescribeImageUsersInput) Validate() error {
//
//	if v.ImageID == nil {
//		return errors.ParameterRequiredError{
//			ParameterName: "ImageID",
//			ParentName:    "DescribeImageUsersInput",
//		}
//	}
//
//	return nil
//}
//
//type DescribeImageUsersOutput struct {
//	Message      *string      `json:"message" name:"message"`
//	Action       *string      `json:"action" name:"action" location:"elements"`
//	ImageUserSet []*ImageUser `json:"image_user_set" name:"image_user_set" location:"elements"`
//	RetCode      *int         `json:"ret_code" name:"ret_code" location:"elements"`
//	TotalCount   *int         `json:"total_count" name:"total_count" location:"elements"`
//}
//

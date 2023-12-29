// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: proto/lpt.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LearningPlanTrackerServiceClient is the client API for LearningPlanTrackerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LearningPlanTrackerServiceClient interface {
	AddCourse(ctx context.Context, in *AddCourseRequest, opts ...grpc.CallOption) (*AddCourseResponse, error)
	AddTopic(ctx context.Context, in *AddTopicRequest, opts ...grpc.CallOption) (*AddTopicResponse, error)
	AddExercise(ctx context.Context, in *AddExerciseRequest, opts ...grpc.CallOption) (*AddExerciseResponse, error)
	ListCourses(ctx context.Context, in *ListCoursesRequest, opts ...grpc.CallOption) (*ListCoursesResponse, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
	GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error)
	ListUsersByRole(ctx context.Context, in *ListUsersByRoleRequest, opts ...grpc.CallOption) (*ListUsersByRoleResponse, error)
	CreateAssignment(ctx context.Context, in *CreateAssignmentRequest, opts ...grpc.CallOption) (*CreateAssignmentResponse, error)
	ListCurrentAssignments(ctx context.Context, in *ListCurrentAssignmentsRequest, opts ...grpc.CallOption) (*ListCurrentAssignmentsResponse, error)
	ListAssignedCourses(ctx context.Context, in *ListAssignedCoursesRequest, opts ...grpc.CallOption) (*ListAssignedCoursesResponse, error)
	GetAssignedCourseDetailsByCourseId(ctx context.Context, in *GetAssignedCourseDetailsByCourseIdRequest, opts ...grpc.CallOption) (*GetAssignedCourseDetailsByCourseIdResponse, error)
	GetAssignedCourseAndMentorDetails(ctx context.Context, in *GetAssignedCourseAndMentorDetailsRequest, opts ...grpc.CallOption) (*GetAssignedCourseAndMentorDetailsResponse, error)
	SubmitExercise(ctx context.Context, in *SubmitExerciseRequest, opts ...grpc.CallOption) (*SubmitExerciseResponse, error)
	DeleteExercise(ctx context.Context, in *DeleteExerciseRequest, opts ...grpc.CallOption) (*DeleteExerciseResponse, error)
	GetSubmittedExercise(ctx context.Context, in *GetSubmittedExerciseRequest, opts ...grpc.CallOption) (*GetSubmittedExerciseResponse, error)
	GetProgress(ctx context.Context, in *GetProgressRequest, opts ...grpc.CallOption) (*GetProgressResponse, error)
	ListAssignedMenteesAndCourses(ctx context.Context, in *ListAssignedMenteesAndCoursesRequest, opts ...grpc.CallOption) (*ListAssignedMenteesAndCoursesResponse, error)
	ListSubmittedExercisesByMentee(ctx context.Context, in *ListSubmittedExercisesByMenteeRequest, opts ...grpc.CallOption) (*ListSubmittedExercisesByMenteeResponse, error)
}

type learningPlanTrackerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLearningPlanTrackerServiceClient(cc grpc.ClientConnInterface) LearningPlanTrackerServiceClient {
	return &learningPlanTrackerServiceClient{cc}
}

func (c *learningPlanTrackerServiceClient) AddCourse(ctx context.Context, in *AddCourseRequest, opts ...grpc.CallOption) (*AddCourseResponse, error) {
	out := new(AddCourseResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/AddCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) AddTopic(ctx context.Context, in *AddTopicRequest, opts ...grpc.CallOption) (*AddTopicResponse, error) {
	out := new(AddTopicResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/AddTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) AddExercise(ctx context.Context, in *AddExerciseRequest, opts ...grpc.CallOption) (*AddExerciseResponse, error) {
	out := new(AddExerciseResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/AddExercise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) ListCourses(ctx context.Context, in *ListCoursesRequest, opts ...grpc.CallOption) (*ListCoursesResponse, error) {
	out := new(ListCoursesResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/ListCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error) {
	out := new(GetUserDetailsResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/GetUserDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) ListUsersByRole(ctx context.Context, in *ListUsersByRoleRequest, opts ...grpc.CallOption) (*ListUsersByRoleResponse, error) {
	out := new(ListUsersByRoleResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/ListUsersByRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) CreateAssignment(ctx context.Context, in *CreateAssignmentRequest, opts ...grpc.CallOption) (*CreateAssignmentResponse, error) {
	out := new(CreateAssignmentResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/CreateAssignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) ListCurrentAssignments(ctx context.Context, in *ListCurrentAssignmentsRequest, opts ...grpc.CallOption) (*ListCurrentAssignmentsResponse, error) {
	out := new(ListCurrentAssignmentsResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/ListCurrentAssignments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) ListAssignedCourses(ctx context.Context, in *ListAssignedCoursesRequest, opts ...grpc.CallOption) (*ListAssignedCoursesResponse, error) {
	out := new(ListAssignedCoursesResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/ListAssignedCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) GetAssignedCourseDetailsByCourseId(ctx context.Context, in *GetAssignedCourseDetailsByCourseIdRequest, opts ...grpc.CallOption) (*GetAssignedCourseDetailsByCourseIdResponse, error) {
	out := new(GetAssignedCourseDetailsByCourseIdResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/GetAssignedCourseDetailsByCourseId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) GetAssignedCourseAndMentorDetails(ctx context.Context, in *GetAssignedCourseAndMentorDetailsRequest, opts ...grpc.CallOption) (*GetAssignedCourseAndMentorDetailsResponse, error) {
	out := new(GetAssignedCourseAndMentorDetailsResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/GetAssignedCourseAndMentorDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) SubmitExercise(ctx context.Context, in *SubmitExerciseRequest, opts ...grpc.CallOption) (*SubmitExerciseResponse, error) {
	out := new(SubmitExerciseResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/SubmitExercise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) DeleteExercise(ctx context.Context, in *DeleteExerciseRequest, opts ...grpc.CallOption) (*DeleteExerciseResponse, error) {
	out := new(DeleteExerciseResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/DeleteExercise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) GetSubmittedExercise(ctx context.Context, in *GetSubmittedExerciseRequest, opts ...grpc.CallOption) (*GetSubmittedExerciseResponse, error) {
	out := new(GetSubmittedExerciseResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/GetSubmittedExercise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) GetProgress(ctx context.Context, in *GetProgressRequest, opts ...grpc.CallOption) (*GetProgressResponse, error) {
	out := new(GetProgressResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/GetProgress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) ListAssignedMenteesAndCourses(ctx context.Context, in *ListAssignedMenteesAndCoursesRequest, opts ...grpc.CallOption) (*ListAssignedMenteesAndCoursesResponse, error) {
	out := new(ListAssignedMenteesAndCoursesResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/ListAssignedMenteesAndCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlanTrackerServiceClient) ListSubmittedExercisesByMentee(ctx context.Context, in *ListSubmittedExercisesByMenteeRequest, opts ...grpc.CallOption) (*ListSubmittedExercisesByMenteeResponse, error) {
	out := new(ListSubmittedExercisesByMenteeResponse)
	err := c.cc.Invoke(ctx, "/LearningPlanTrackerService/ListSubmittedExercisesByMentee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LearningPlanTrackerServiceServer is the server API for LearningPlanTrackerService service.
// All implementations must embed UnimplementedLearningPlanTrackerServiceServer
// for forward compatibility
type LearningPlanTrackerServiceServer interface {
	AddCourse(context.Context, *AddCourseRequest) (*AddCourseResponse, error)
	AddTopic(context.Context, *AddTopicRequest) (*AddTopicResponse, error)
	AddExercise(context.Context, *AddExerciseRequest) (*AddExerciseResponse, error)
	ListCourses(context.Context, *ListCoursesRequest) (*ListCoursesResponse, error)
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
	GetUserDetails(context.Context, *GetUserDetailsRequest) (*GetUserDetailsResponse, error)
	ListUsersByRole(context.Context, *ListUsersByRoleRequest) (*ListUsersByRoleResponse, error)
	CreateAssignment(context.Context, *CreateAssignmentRequest) (*CreateAssignmentResponse, error)
	ListCurrentAssignments(context.Context, *ListCurrentAssignmentsRequest) (*ListCurrentAssignmentsResponse, error)
	ListAssignedCourses(context.Context, *ListAssignedCoursesRequest) (*ListAssignedCoursesResponse, error)
	GetAssignedCourseDetailsByCourseId(context.Context, *GetAssignedCourseDetailsByCourseIdRequest) (*GetAssignedCourseDetailsByCourseIdResponse, error)
	GetAssignedCourseAndMentorDetails(context.Context, *GetAssignedCourseAndMentorDetailsRequest) (*GetAssignedCourseAndMentorDetailsResponse, error)
	SubmitExercise(context.Context, *SubmitExerciseRequest) (*SubmitExerciseResponse, error)
	DeleteExercise(context.Context, *DeleteExerciseRequest) (*DeleteExerciseResponse, error)
	GetSubmittedExercise(context.Context, *GetSubmittedExerciseRequest) (*GetSubmittedExerciseResponse, error)
	GetProgress(context.Context, *GetProgressRequest) (*GetProgressResponse, error)
	ListAssignedMenteesAndCourses(context.Context, *ListAssignedMenteesAndCoursesRequest) (*ListAssignedMenteesAndCoursesResponse, error)
	ListSubmittedExercisesByMentee(context.Context, *ListSubmittedExercisesByMenteeRequest) (*ListSubmittedExercisesByMenteeResponse, error)
	mustEmbedUnimplementedLearningPlanTrackerServiceServer()
}

// UnimplementedLearningPlanTrackerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLearningPlanTrackerServiceServer struct {
}

func (UnimplementedLearningPlanTrackerServiceServer) AddCourse(context.Context, *AddCourseRequest) (*AddCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCourse not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) AddTopic(context.Context, *AddTopicRequest) (*AddTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTopic not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) AddExercise(context.Context, *AddExerciseRequest) (*AddExerciseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExercise not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) ListCourses(context.Context, *ListCoursesRequest) (*ListCoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourses not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) GetUserDetails(context.Context, *GetUserDetailsRequest) (*GetUserDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetails not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) ListUsersByRole(context.Context, *ListUsersByRoleRequest) (*ListUsersByRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsersByRole not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) CreateAssignment(context.Context, *CreateAssignmentRequest) (*CreateAssignmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAssignment not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) ListCurrentAssignments(context.Context, *ListCurrentAssignmentsRequest) (*ListCurrentAssignmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCurrentAssignments not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) ListAssignedCourses(context.Context, *ListAssignedCoursesRequest) (*ListAssignedCoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAssignedCourses not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) GetAssignedCourseDetailsByCourseId(context.Context, *GetAssignedCourseDetailsByCourseIdRequest) (*GetAssignedCourseDetailsByCourseIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssignedCourseDetailsByCourseId not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) GetAssignedCourseAndMentorDetails(context.Context, *GetAssignedCourseAndMentorDetailsRequest) (*GetAssignedCourseAndMentorDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssignedCourseAndMentorDetails not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) SubmitExercise(context.Context, *SubmitExerciseRequest) (*SubmitExerciseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitExercise not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) DeleteExercise(context.Context, *DeleteExerciseRequest) (*DeleteExerciseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExercise not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) GetSubmittedExercise(context.Context, *GetSubmittedExerciseRequest) (*GetSubmittedExerciseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmittedExercise not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) GetProgress(context.Context, *GetProgressRequest) (*GetProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProgress not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) ListAssignedMenteesAndCourses(context.Context, *ListAssignedMenteesAndCoursesRequest) (*ListAssignedMenteesAndCoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAssignedMenteesAndCourses not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) ListSubmittedExercisesByMentee(context.Context, *ListSubmittedExercisesByMenteeRequest) (*ListSubmittedExercisesByMenteeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubmittedExercisesByMentee not implemented")
}
func (UnimplementedLearningPlanTrackerServiceServer) mustEmbedUnimplementedLearningPlanTrackerServiceServer() {
}

// UnsafeLearningPlanTrackerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LearningPlanTrackerServiceServer will
// result in compilation errors.
type UnsafeLearningPlanTrackerServiceServer interface {
	mustEmbedUnimplementedLearningPlanTrackerServiceServer()
}

func RegisterLearningPlanTrackerServiceServer(s grpc.ServiceRegistrar, srv LearningPlanTrackerServiceServer) {
	s.RegisterService(&LearningPlanTrackerService_ServiceDesc, srv)
}

func _LearningPlanTrackerService_AddCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).AddCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/AddCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).AddCourse(ctx, req.(*AddCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_AddTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).AddTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/AddTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).AddTopic(ctx, req.(*AddTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_AddExercise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).AddExercise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/AddExercise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).AddExercise(ctx, req.(*AddExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_ListCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCoursesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).ListCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/ListCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).ListCourses(ctx, req.(*ListCoursesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_GetUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).GetUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/GetUserDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).GetUserDetails(ctx, req.(*GetUserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_ListUsersByRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersByRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).ListUsersByRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/ListUsersByRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).ListUsersByRole(ctx, req.(*ListUsersByRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_CreateAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAssignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).CreateAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/CreateAssignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).CreateAssignment(ctx, req.(*CreateAssignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_ListCurrentAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCurrentAssignmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).ListCurrentAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/ListCurrentAssignments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).ListCurrentAssignments(ctx, req.(*ListCurrentAssignmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_ListAssignedCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAssignedCoursesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).ListAssignedCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/ListAssignedCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).ListAssignedCourses(ctx, req.(*ListAssignedCoursesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_GetAssignedCourseDetailsByCourseId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssignedCourseDetailsByCourseIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).GetAssignedCourseDetailsByCourseId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/GetAssignedCourseDetailsByCourseId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).GetAssignedCourseDetailsByCourseId(ctx, req.(*GetAssignedCourseDetailsByCourseIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_GetAssignedCourseAndMentorDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssignedCourseAndMentorDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).GetAssignedCourseAndMentorDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/GetAssignedCourseAndMentorDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).GetAssignedCourseAndMentorDetails(ctx, req.(*GetAssignedCourseAndMentorDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_SubmitExercise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).SubmitExercise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/SubmitExercise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).SubmitExercise(ctx, req.(*SubmitExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_DeleteExercise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).DeleteExercise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/DeleteExercise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).DeleteExercise(ctx, req.(*DeleteExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_GetSubmittedExercise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmittedExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).GetSubmittedExercise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/GetSubmittedExercise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).GetSubmittedExercise(ctx, req.(*GetSubmittedExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_GetProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).GetProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/GetProgress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).GetProgress(ctx, req.(*GetProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_ListAssignedMenteesAndCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAssignedMenteesAndCoursesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).ListAssignedMenteesAndCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/ListAssignedMenteesAndCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).ListAssignedMenteesAndCourses(ctx, req.(*ListAssignedMenteesAndCoursesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlanTrackerService_ListSubmittedExercisesByMentee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubmittedExercisesByMenteeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlanTrackerServiceServer).ListSubmittedExercisesByMentee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LearningPlanTrackerService/ListSubmittedExercisesByMentee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlanTrackerServiceServer).ListSubmittedExercisesByMentee(ctx, req.(*ListSubmittedExercisesByMenteeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LearningPlanTrackerService_ServiceDesc is the grpc.ServiceDesc for LearningPlanTrackerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LearningPlanTrackerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LearningPlanTrackerService",
	HandlerType: (*LearningPlanTrackerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCourse",
			Handler:    _LearningPlanTrackerService_AddCourse_Handler,
		},
		{
			MethodName: "AddTopic",
			Handler:    _LearningPlanTrackerService_AddTopic_Handler,
		},
		{
			MethodName: "AddExercise",
			Handler:    _LearningPlanTrackerService_AddExercise_Handler,
		},
		{
			MethodName: "ListCourses",
			Handler:    _LearningPlanTrackerService_ListCourses_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _LearningPlanTrackerService_AddUser_Handler,
		},
		{
			MethodName: "GetUserDetails",
			Handler:    _LearningPlanTrackerService_GetUserDetails_Handler,
		},
		{
			MethodName: "ListUsersByRole",
			Handler:    _LearningPlanTrackerService_ListUsersByRole_Handler,
		},
		{
			MethodName: "CreateAssignment",
			Handler:    _LearningPlanTrackerService_CreateAssignment_Handler,
		},
		{
			MethodName: "ListCurrentAssignments",
			Handler:    _LearningPlanTrackerService_ListCurrentAssignments_Handler,
		},
		{
			MethodName: "ListAssignedCourses",
			Handler:    _LearningPlanTrackerService_ListAssignedCourses_Handler,
		},
		{
			MethodName: "GetAssignedCourseDetailsByCourseId",
			Handler:    _LearningPlanTrackerService_GetAssignedCourseDetailsByCourseId_Handler,
		},
		{
			MethodName: "GetAssignedCourseAndMentorDetails",
			Handler:    _LearningPlanTrackerService_GetAssignedCourseAndMentorDetails_Handler,
		},
		{
			MethodName: "SubmitExercise",
			Handler:    _LearningPlanTrackerService_SubmitExercise_Handler,
		},
		{
			MethodName: "DeleteExercise",
			Handler:    _LearningPlanTrackerService_DeleteExercise_Handler,
		},
		{
			MethodName: "GetSubmittedExercise",
			Handler:    _LearningPlanTrackerService_GetSubmittedExercise_Handler,
		},
		{
			MethodName: "GetProgress",
			Handler:    _LearningPlanTrackerService_GetProgress_Handler,
		},
		{
			MethodName: "ListAssignedMenteesAndCourses",
			Handler:    _LearningPlanTrackerService_ListAssignedMenteesAndCourses_Handler,
		},
		{
			MethodName: "ListSubmittedExercisesByMentee",
			Handler:    _LearningPlanTrackerService_ListSubmittedExercisesByMentee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/lpt.proto",
}

# \DefaultApi

All URIs are relative to *https://api.example.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CourseCourseIdDelete**](DefaultApi.md#CourseCourseIdDelete) | **Delete** /course/{courseId} | Delete specific course
[**CourseCourseIdGet**](DefaultApi.md#CourseCourseIdGet) | **Get** /course/{courseId} | Return specific course
[**CourseCourseIdPost**](DefaultApi.md#CourseCourseIdPost) | **Post** /course/{courseId} | Updates an existing course
[**CourseGet**](DefaultApi.md#CourseGet) | **Get** /course | Returns a list of users.
[**CoursePost**](DefaultApi.md#CoursePost) | **Post** /course | Creates a new course
[**StudentGet**](DefaultApi.md#StudentGet) | **Get** /student | Returns a list of students.
[**StudentPost**](DefaultApi.md#StudentPost) | **Post** /student | Creates a new student
[**StudentStudentIdGet**](DefaultApi.md#StudentStudentIdGet) | **Get** /student/{studentId} | Return specific student
[**StudentStudentIdPost**](DefaultApi.md#StudentStudentIdPost) | **Post** /student/{studentId} | Updates an existing student
[**TeacherGet**](DefaultApi.md#TeacherGet) | **Get** /teacher | Returns a list of users.
[**TeacherPost**](DefaultApi.md#TeacherPost) | **Post** /teacher | Creates a new teacher
[**TeacherTeacherIdGet**](DefaultApi.md#TeacherTeacherIdGet) | **Get** /teacher/{teacherId} | Return specific teacher
[**TeacherTeacherIdPost**](DefaultApi.md#TeacherTeacherIdPost) | **Post** /teacher/{teacherId} | Updates an existing teacher


# **CourseCourseIdDelete**
> CourseCourseIdDelete(ctx, courseId)
Delete specific course

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **courseId** | **int32**| Course ID to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CourseCourseIdGet**
> Course CourseCourseIdGet(ctx, courseId)
Return specific course

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **courseId** | **int32**| ID of course to return | 

### Return type

[**Course**](course.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CourseCourseIdPost**
> CourseCourseIdPost(ctx, courseId, optional)
Updates an existing course

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **courseId** | **int32**| Course that needs to be updated | 
 **optional** | ***DefaultApiCourseCourseIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiCourseCourseIdPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| To update the course name | 
 **teacher** | **optional.Int32**| To update the course teacher | 
 **studentSatisfactionRating** | **optional.Int32**| To update the course Satisfaction Rating | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CourseGet**
> Course CourseGet(ctx, )
Returns a list of users.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Course**](course.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CoursePost**
> CoursePost(ctx, addCourse)
Creates a new course

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addCourse** | [**Course**](Course.md)| Creates a new course with courseid, coursename and courseteacher | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StudentGet**
> Student StudentGet(ctx, )
Returns a list of students.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Student**](student.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StudentPost**
> StudentPost(ctx, addStudent)
Creates a new student

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addStudent** | [**Student**](Student.md)| Creates a new student with id, name and workload | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StudentStudentIdGet**
> Student StudentStudentIdGet(ctx, studentId)
Return specific student

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **studentId** | **int32**| ID of student to return | 

### Return type

[**Student**](student.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StudentStudentIdPost**
> StudentStudentIdPost(ctx, studentId, optional)
Updates an existing student

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **studentId** | **int32**| Student who needs to be updated | 
 **optional** | ***DefaultApiStudentStudentIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiStudentStudentIdPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| To update the students name | 
 **workload** | **optional.Int32**| To update the students workload | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeacherGet**
> Teacher TeacherGet(ctx, )
Returns a list of users.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Teacher**](teacher.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeacherPost**
> TeacherPost(ctx, addTeacher)
Creates a new teacher

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addTeacher** | [**Teacher**](Teacher.md)| Creates a new teacher with teacherid and teacherename | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeacherTeacherIdGet**
> Teacher TeacherTeacherIdGet(ctx, teacherId)
Return specific teacher

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teacherId** | **int32**| ID of teacher to return | 

### Return type

[**Teacher**](teacher.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeacherTeacherIdPost**
> TeacherTeacherIdPost(ctx, teacherId, optional)
Updates an existing teacher

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teacherId** | **int32**| Teacher who needs to be updated | 
 **optional** | ***DefaultApiTeacherTeacherIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiTeacherTeacherIdPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| To update the teachers name | 
 **popularityScore** | **optional.Int32**| To update the teacher Popularity Score | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


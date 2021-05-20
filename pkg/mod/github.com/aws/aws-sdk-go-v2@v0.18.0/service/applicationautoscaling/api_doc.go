// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package applicationautoscaling provides the client and types for making API
// requests to Application Auto Scaling.
//
// With Application Auto Scaling, you can configure automatic scaling for the
// following resources:
//
//    * Amazon ECS services
//
//    * Amazon EC2 Spot Fleet requests
//
//    * Amazon EMR clusters
//
//    * Amazon AppStream 2.0 fleets
//
//    * Amazon DynamoDB tables and global secondary indexes throughput capacity
//
//    * Amazon Aurora Replicas
//
//    * Amazon SageMaker endpoint variants
//
//    * Custom resources provided by your own applications or services
//
//    * Amazon Comprehend document classification endpoints
//
//    * AWS Lambda function provisioned concurrency
//
// API Summary
//
// The Application Auto Scaling service API includes three key sets of actions:
//
//    * Register and manage scalable targets - Register AWS or custom resources
//    as scalable targets (a resource that Application Auto Scaling can scale),
//    set minimum and maximum capacity limits, and retrieve information on existing
//    scalable targets.
//
//    * Configure and manage automatic scaling - Define scaling policies to
//    dynamically scale your resources in response to CloudWatch alarms, schedule
//    one-time or recurring scaling actions, and retrieve your recent scaling
//    activity history.
//
//    * Suspend and resume scaling - Temporarily suspend and later resume automatic
//    scaling by calling the RegisterScalableTarget action for any Application
//    Auto Scaling scalable target. You can suspend and resume, individually
//    or in combination, scale-out activities triggered by a scaling policy,
//    scale-in activities triggered by a scaling policy, and scheduled scaling.
//
// To learn more about Application Auto Scaling, including information about
// granting IAM users required permissions for Application Auto Scaling actions,
// see the Application Auto Scaling User Guide (https://docs.aws.amazon.com/autoscaling/application/userguide/what-is-application-auto-scaling.html).
//
// See https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06 for more information on this service.
//
// See applicationautoscaling package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/applicationautoscaling/
//
// Using the Client
//
// To use Application Auto Scaling with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Application Auto Scaling client for more information on
// creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/applicationautoscaling/#New
package applicationautoscaling

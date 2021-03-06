// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ses

const (

	// ErrCodeAlreadyExistsException for service response error code
	// "AlreadyExists".
	//
	// Indicates that a resource could not be created because of a naming conflict.
	ErrCodeAlreadyExistsException = "AlreadyExists"

	// ErrCodeCannotDeleteException for service response error code
	// "CannotDelete".
	//
	// Indicates that the delete operation could not be completed.
	ErrCodeCannotDeleteException = "CannotDelete"

	// ErrCodeConfigurationSetAlreadyExistsException for service response error code
	// "ConfigurationSetAlreadyExists".
	//
	// Indicates that the configuration set could not be created because of a naming
	// conflict.
	ErrCodeConfigurationSetAlreadyExistsException = "ConfigurationSetAlreadyExists"

	// ErrCodeConfigurationSetDoesNotExistException for service response error code
	// "ConfigurationSetDoesNotExist".
	//
	// Indicates that the configuration set does not exist.
	ErrCodeConfigurationSetDoesNotExistException = "ConfigurationSetDoesNotExist"

	// ErrCodeEventDestinationAlreadyExistsException for service response error code
	// "EventDestinationAlreadyExists".
	//
	// Indicates that the event destination could not be created because of a naming
	// conflict.
	ErrCodeEventDestinationAlreadyExistsException = "EventDestinationAlreadyExists"

	// ErrCodeEventDestinationDoesNotExistException for service response error code
	// "EventDestinationDoesNotExist".
	//
	// Indicates that the event destination does not exist.
	ErrCodeEventDestinationDoesNotExistException = "EventDestinationDoesNotExist"

	// ErrCodeInvalidCloudWatchDestinationException for service response error code
	// "InvalidCloudWatchDestination".
	//
	// Indicates that the Amazon CloudWatch destination is invalid. See the error
	// message for details.
	ErrCodeInvalidCloudWatchDestinationException = "InvalidCloudWatchDestination"

	// ErrCodeInvalidConfigurationSetException for service response error code
	// "InvalidConfigurationSet".
	//
	// Indicates that the configuration set is invalid. See the error message for
	// details.
	ErrCodeInvalidConfigurationSetException = "InvalidConfigurationSet"

	// ErrCodeInvalidFirehoseDestinationException for service response error code
	// "InvalidFirehoseDestination".
	//
	// Indicates that the Amazon Kinesis Firehose destination is invalid. See the
	// error message for details.
	ErrCodeInvalidFirehoseDestinationException = "InvalidFirehoseDestination"

	// ErrCodeInvalidLambdaFunctionException for service response error code
	// "InvalidLambdaFunction".
	//
	// Indicates that the provided AWS Lambda function is invalid, or that Amazon
	// SES could not execute the provided function, possibly due to permissions
	// issues. For information about giving permissions, see the Amazon SES Developer
	// Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html).
	ErrCodeInvalidLambdaFunctionException = "InvalidLambdaFunction"

	// ErrCodeInvalidPolicyException for service response error code
	// "InvalidPolicy".
	//
	// Indicates that the provided policy is invalid. Check the error stack for
	// more information about what caused the error.
	ErrCodeInvalidPolicyException = "InvalidPolicy"

	// ErrCodeInvalidS3ConfigurationException for service response error code
	// "InvalidS3Configuration".
	//
	// Indicates that the provided Amazon S3 bucket or AWS KMS encryption key is
	// invalid, or that Amazon SES could not publish to the bucket, possibly due
	// to permissions issues. For information about giving permissions, see the
	// Amazon SES Developer Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html).
	ErrCodeInvalidS3ConfigurationException = "InvalidS3Configuration"

	// ErrCodeInvalidSnsTopicException for service response error code
	// "InvalidSnsTopic".
	//
	// Indicates that the provided Amazon SNS topic is invalid, or that Amazon SES
	// could not publish to the topic, possibly due to permissions issues. For information
	// about giving permissions, see the Amazon SES Developer Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html).
	ErrCodeInvalidSnsTopicException = "InvalidSnsTopic"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceeded".
	//
	// Indicates that a resource could not be created because of service limits.
	// For a list of Amazon SES limits, see the Amazon SES Developer Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/limits.html).
	ErrCodeLimitExceededException = "LimitExceeded"

	// ErrCodeMailFromDomainNotVerifiedException for service response error code
	// "MailFromDomainNotVerifiedException".
	//
	// Indicates that the message could not be sent because Amazon SES could not
	// read the MX record required to use the specified MAIL FROM domain. For information
	// about editing the custom MAIL FROM domain settings for an identity, see the
	// Amazon SES Developer Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/mail-from-edit.html).
	ErrCodeMailFromDomainNotVerifiedException = "MailFromDomainNotVerifiedException"

	// ErrCodeMessageRejected for service response error code
	// "MessageRejected".
	//
	// Indicates that the action failed, and the message could not be sent. Check
	// the error stack for more information about what caused the error.
	ErrCodeMessageRejected = "MessageRejected"

	// ErrCodeRuleDoesNotExistException for service response error code
	// "RuleDoesNotExist".
	//
	// Indicates that the provided receipt rule does not exist.
	ErrCodeRuleDoesNotExistException = "RuleDoesNotExist"

	// ErrCodeRuleSetDoesNotExistException for service response error code
	// "RuleSetDoesNotExist".
	//
	// Indicates that the provided receipt rule set does not exist.
	ErrCodeRuleSetDoesNotExistException = "RuleSetDoesNotExist"
)

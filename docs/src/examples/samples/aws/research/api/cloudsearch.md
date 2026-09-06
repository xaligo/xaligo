# Amazon CloudSearch

API version: 2013-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cloudsearch/2013-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BuildSuggesters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldNames` | `List<string>` | no |

## CreateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainStatus` | `DomainStatus` | no |

## DefineAnalysisScheme

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `AnalysisScheme` | `AnalysisScheme` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisScheme` | `AnalysisSchemeStatus` | yes |

## DefineExpression

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Expression` | `Expression` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Expression` | `ExpressionStatus` | yes |

## DefineIndexField

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `IndexField` | `IndexField` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexField` | `IndexFieldStatus` | yes |

## DefineSuggester

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Suggester` | `Suggester` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Suggester` | `SuggesterStatus` | yes |

## DeleteAnalysisScheme

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `AnalysisSchemeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisScheme` | `AnalysisSchemeStatus` | yes |

## DeleteDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainStatus` | `DomainStatus` | no |

## DeleteExpression

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ExpressionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Expression` | `ExpressionStatus` | yes |

## DeleteIndexField

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `IndexFieldName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexField` | `IndexFieldStatus` | yes |

## DeleteSuggester

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SuggesterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Suggester` | `SuggesterStatus` | yes |

## DescribeAnalysisSchemes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `AnalysisSchemeNames` | `List<string>` | no |
| `Deployed` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisSchemes` | `List<AnalysisSchemeStatus>` | yes |

## DescribeAvailabilityOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Deployed` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityOptions` | `AvailabilityOptionsStatus` | no |

## DescribeDomainEndpointOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Deployed` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainEndpointOptions` | `DomainEndpointOptionsStatus` | no |

## DescribeDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainStatusList` | `List<DomainStatus>` | yes |

## DescribeExpressions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ExpressionNames` | `List<string>` | no |
| `Deployed` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Expressions` | `List<ExpressionStatus>` | yes |

## DescribeIndexFields

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `FieldNames` | `List<string>` | no |
| `Deployed` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexFields` | `List<IndexFieldStatus>` | yes |

## DescribeScalingParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingParameters` | `ScalingParametersStatus` | yes |

## DescribeServiceAccessPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Deployed` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPolicies` | `AccessPoliciesStatus` | yes |

## DescribeSuggesters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SuggesterNames` | `List<string>` | no |
| `Deployed` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Suggesters` | `List<SuggesterStatus>` | yes |

## IndexDocuments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldNames` | `List<string>` | no |

## ListDomainNames

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainNames` | `Map<string>` | no |

## UpdateAvailabilityOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MultiAZ` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityOptions` | `AvailabilityOptionsStatus` | no |

## UpdateDomainEndpointOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DomainEndpointOptions` | `DomainEndpointOptions` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainEndpointOptions` | `DomainEndpointOptionsStatus` | no |

## UpdateScalingParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ScalingParameters` | `ScalingParameters` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingParameters` | `ScalingParametersStatus` | yes |

## UpdateServiceAccessPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `AccessPolicies` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPolicies` | `AccessPoliciesStatus` | yes |


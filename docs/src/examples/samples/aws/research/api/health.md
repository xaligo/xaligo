# AWS Health APIs and Notifications

API version: 2016-08-04. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/health/2016-08-04/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DescribeAffectedAccountsForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `affectedAccounts` | `List<string>` | no |
| `eventScopeCode` | `string` | no |
| `nextToken` | `string` | no |

## DescribeAffectedEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `EntityFilter` | yes |
| `locale` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entities` | `List<AffectedEntity>` | no |
| `nextToken` | `string` | no |

## DescribeAffectedEntitiesForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `organizationEntityFilters` | `List<EventAccountFilter>` | no |
| `locale` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `organizationEntityAccountFilters` | `List<EntityAccountFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entities` | `List<AffectedEntity>` | no |
| `failedSet` | `List<OrganizationAffectedEntitiesErrorItem>` | no |
| `nextToken` | `string` | no |

## DescribeEntityAggregates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventArns` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entityAggregates` | `List<EntityAggregate>` | no |

## DescribeEntityAggregatesForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventArns` | `List<string>` | yes |
| `awsAccountIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `organizationEntityAggregates` | `List<OrganizationEntityAggregate>` | no |

## DescribeEventAggregates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `EventFilter` | no |
| `aggregateField` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventAggregates` | `List<EventAggregate>` | no |
| `nextToken` | `string` | no |

## DescribeEventDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventArns` | `List<string>` | yes |
| `locale` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulSet` | `List<EventDetails>` | no |
| `failedSet` | `List<EventDetailsErrorItem>` | no |

## DescribeEventDetailsForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `organizationEventDetailFilters` | `List<EventAccountFilter>` | yes |
| `locale` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulSet` | `List<OrganizationEventDetails>` | no |
| `failedSet` | `List<OrganizationEventDetailsErrorItem>` | no |

## DescribeEventTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `EventTypeFilter` | no |
| `locale` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventTypes` | `List<EventType>` | no |
| `nextToken` | `string` | no |

## DescribeEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `EventFilter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `locale` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<Event>` | no |
| `nextToken` | `string` | no |

## DescribeEventsForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `OrganizationEventFilter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `locale` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<OrganizationEvent>` | no |
| `nextToken` | `string` | no |

## DescribeHealthServiceStatusForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `healthServiceAccessStatusForOrganization` | `string` | no |

## DisableHealthServiceAccessForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableHealthServiceAccessForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|



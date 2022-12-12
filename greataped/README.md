[![image](https://img.shields.io/badge/Go-00A7D0?style=for-the-badge&logo=go&logoColor=white)](https://go.dev) [![image](https://img.shields.io/badge/ActivityPub-DD307D?style=for-the-badge&logoColor=white)](https://www.w3.org/TR/activitypub/) [![image](https://img.shields.io/badge/JSON--LD-FF6600?style=for-the-badge&logo=json&logoColor=white)](https://json-ld.org) [![image](https://img.shields.io/badge/MySQL-32738C?style=for-the-badge&logo=mysql&logoColor=white)](https://www.mysql.com) [![image](https://img.shields.io/badge/MariaDB-39818D?style=for-the-badge&logo=mariadb&logoColor=white)](https://mariadb.com)

## Note

⚠️ This project is under heavy development and should not be used in production yet.

## APIs:
1. [Echo](#echo)
2. [Signup](#signup)
3. [Verify](#verify)
4. [Login](#login)
5. [GetProfileByUser](#get-profile-by-user)
6. [UpdateProfileByUser](#update-profile-by-user)
7. [Logout](#logout)
8. [Webfinger](#webfinger)
9. [GetActor](#get-actor)
10. [FollowActor](#follow-actor)
11. [AuthorizeInteraction](#authorize-interaction)
12. [GetFollowers](#get-followers)
13. [GetFollowing](#get-following)

---

## Echo
```
Request:
    Document document

Result:
    Document document
```
[Back to List](#apis)

## Signup
```
Request:
    string username
    string email
    // Should be at least 7 characters including upper and lowercase, digits and symbols
    string password

Result:
    string token
    string code
```
[Back to List](#apis)

## Verify
```
Request:
    string email
    string token
    string code

Result:
    string token
```
[Back to List](#apis)

## Login
```
Request:
    string email
    string password

Result:
    string username
    string token
```
[Back to List](#apis)

## Get Profile By User
```
Request:

Result:
    string username
    string displayName
    string avatar
    string banner
    string summary
    string github
```
[Back to List](#apis)

## Update Profile By User
```
Request:
    string displayName
    string avatar
    string banner
    string summary
    string github

Result:
    string displayName
    string avatar
    string banner
    string summary
    string github
```
[Back to List](#apis)

## Logout
```
Request:

Result:
```
[Back to List](#apis)

## Webfinger
```
Request:
    string resource

Result:
    repeated string aliases
    repeated ActivityPubLink links
    string subject
```
[Back to List](#apis)

## Get Actor
```
Request:
    string username

Result:
    repeated string @context
    string id
    string followers
    string following
    string inbox
    string outbox
    string name
    string preferredUsername
    string type
    string url
    ActivityPubMedia icon
    ActivityPubMedia image
    ActivityPubPublicKey publicKey
    string summary
    string published
```
[Back to List](#apis)

## Follow Actor
```
Request:
    string username
    string acct

Result:
    string url
```
[Back to List](#apis)

## Authorize Interaction
```
Request:
    string uri

Result:
    string uri
    bool success
```
[Back to List](#apis)

## Get Followers
```
Request:
    string username

Result:
    string @context
    string id
    string type
    int32 totalItems
    repeated string orderedItems
    string first
```
[Back to List](#apis)

## Get Following
```
Request:
    string username

Result:
    string @context
    string id
    string type
    int32 totalItems
    repeated string orderedItems
    string first
```
[Back to List](#apis)

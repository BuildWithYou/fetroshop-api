Coding Preference
=================

### Naming Convention
1. API request : camelCase
2. API response : camelCase
3. Database column : snake_case
4. Variable : camelCase
5. Type name separator : _
6. File name separator : -

### Coding Style
1. Function: returning pointer

### Data Processing Rules
1. Identifier: ID used inside system, username/code used outside system

### Responses
1. After create some entity:
   - Send code 201
   - Send status 'Created'
   - Send message '{entity} created successfully'
   - Send detail entity in 'data' field
0. After update some entity:
   - Send code 200
   - Send status 'OK'
   - Send message '{entity} updated successfully'
   - Send detail entity in 'data' field
0. If the message is single string:
   - Send constant string in 'message' field
   - Send detail message in 'meta.message'
0. If the message is array of strings:
   - Send constant string in 'message' field
   - Send detail message in 'meta.messages'


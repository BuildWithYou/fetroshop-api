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

### Error response
1. If the message is single string:
   - Send constant string in 'message' field
   - Send detail message in 'meta.message'
2. If the message is array of strings:
   - Send constant string in 'message' field
   - Send detail message in 'meta.messages'


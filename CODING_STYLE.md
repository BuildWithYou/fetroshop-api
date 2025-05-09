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
1. Identifier: `ID` used inside system, `username`/`code` used outside system

### Responses
1. After create some entity:
   - Send `code` 201
   - Send `status` 'Created'
   - Send `message` '{entity} created successfully'
   - Send detail entity in `data` field
2. After update some entity:
   - Send `code` 200
   - Send `status` 'OK'
   - Send `message` '{entity} updated successfully'
   - Send detail entity in `data` field
3. If the message is single string:
   - Send constant string in `message` field
   - Send detail message in '`meta.message`'
4. If the message is array of strings:
   - Send constant string in `message` field
   - Send detail message in `meta.messages`

### Error Handling
1. If error is system error:
   - Return error and let it handled by default error handler if need to log the error
   - Return response 500 if don't need to log the error
2. If error triggered by user input, return `message` "VALIDATION ERROR" and put the errors in `errors` field

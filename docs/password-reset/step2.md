# Password Reset - Step 2: Verification & Password Update

## 1. Step Summary
Allow users to complete password reset using verification link with new password.

## 2. Inputs
- Password reset token (from URL/link)
- New password (user input)
- Password confirmation (must match)

## 3. Dependencies
- Database record of the reset token (created in [Step 1: Request](step1.md))
- User database record
- Password hashing service

## 4. Behavior Scenarios

**I. Valid token and passwords match**
1. User clicks verification link with valid token
2. User enters matching new passwords
3. System validates token and updates password
4. Token is deleted (single-use)
5. User redirected to success page

**II. Token expired**
1. User clicks link with expired token
2. System shows "token expired" error
3. User can request new reset

**III. Passwords don't match**
1. User enters non-matching passwords
2. System shows validation error
3. User can retry without new token

## 5. System Constraints
- Passwords must meet security requirements
- Tokens valid for 1 hour only
- Operation must complete within 5 seconds
- Password update and token deletion must be atomic

## 6. Side Effects
- User password hash updated in database
- Reset token removed from database
- Password change logged in audit trail
- User may need to re-login with new password
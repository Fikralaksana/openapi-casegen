# Password Reset - Step 1: Request

## 1. Step Summary

Allow users to initiate password reset by submitting their email address. The system validates the email, generates a secure reset token, and sends a password reset email with a verification link.

## 2. Inputs

- User's email address (for password reset request)

## 3. Dependencies

- Database record of a user
- Email service
- Database record of the token
- Identity and access management (keycloak)

## 4. Behavior Scenarios

**I. Reset request success**
1. User submits valid email address for password reset
2. Generate secure reset token and send email with reset link
3. User receives password reset email

**II. Reset request invalid email**
1. User submits non-existent email address
2. Return success response (security by obscurity)
3. No email sent, no error exposed to potential attacker

## 5. System Constraints

- Reset tokens must expire within 1 hour
- Maximum 5 reset attempts per hour per email
- Emails must be sent within 30 seconds of request
- Token links must be single-use only

## 6. Side Effects

- New reset password token will be added to the existing tokens table when reset attempt succeeds
- Email notification will be sent to user when password reset is initiated (links to [Step 2: Verification](step2.md))
- Security monitoring may flag unusual reset patterns
- User engagement metrics may improve due to self-service recovery
- Existing authentication flows remain unaffected
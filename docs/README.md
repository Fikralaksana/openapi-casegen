# Feature Kickoff Notes

This directory contains structured notes taken during feature kickoff meetings. Each feature has a dedicated folder with kickoff documentation that captures business requirements, technical constraints, and implementation decisions made during the design phase.

## Feature Structure

Features are organized in dedicated folders with the following structure:

```
feature-name/
├── flow.md        # Feature overview and flow description
├── step1.md       # First implementation step
├── step2.md       # Second implementation step
└── ...           # Additional steps as needed
```

## Example: Password Reset

**[`password-reset/`](password-reset/)** - Complete example of a multi-step feature

- **[`flow.md`](password-reset/flow.md)** - Feature overview and high-level flow
- **[`step1.md`](password-reset/step1.md)** - Password reset request implementation
- **[`step2.md`](password-reset/step2.md)** - Password verification and update implementation

## Template Structure

The password reset feature serves as a template for kickoff notes. Copy the `password-reset/` folder structure for new features and adapt the content to capture the specific decisions and requirements discussed during your kickoff meetings.

## How to Use

1. **Create during kickoff** - Take structured notes while discussing feature requirements with stakeholders
2. **Capture decisions** - Document business rules, constraints, and acceptance criteria agreed upon
3. **Break down complexity** - Split complex features into digestible steps for discussion
4. **Reference during development** - Use these notes as the source of truth for implementation
5. **Validate completion** - Ensure all kickoff decisions are properly implemented

## Benefits

- **Structured kickoff discussions** - Ensures all aspects are covered during design
- **Clear decision documentation** - Captures agreements made during meetings
- **Implementation guidance** - Provides clear requirements for development teams
- **Consistency across features** - Standardized format for all feature kickoffs
- **Audit trail** - Historical record of design decisions and rationale

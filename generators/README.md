# Test Case Generators

This package contains modular test case generators for different OpenAPI data types. Each data type has its own generator that can be easily extended or modified.

## Structure

- `base.go` - Common interfaces and main dispatcher
- `enum.go` - Enum parameter test cases
- `integer.go` - Integer parameter test cases
- `number.go` - Number (float) parameter test cases
- `string.go` - String parameter test cases
- `boolean.go` - Boolean parameter test cases

## Adding a New Data Type

1. Create a new file `datatype.go` (e.g., `array.go`, `object.go`)
2. Implement the `Generator` interface:

```go
type YourTypeGenerator struct{}

func (g *YourTypeGenerator) GenerateTestCases(baseID string, paramType string, enumValues []interface{}) []TestCase {
    return []TestCase{
        // Your test cases here
    }
}
```

3. Add the new type to the switch statement in `base.go`:

```go
case "your_type":
    generator = &YourTypeGenerator{}
```

## Test Case Types

- `valid` - Valid inputs for the data type
- `invalid` - Invalid inputs that should cause errors
- `boundary_min` - Minimum boundary values
- `boundary_max` - Maximum boundary values
- `enum_value` - Individual enum values (for enum parameters)

## Example Output

For an integer parameter `limit`:
- `users.limit_valid_input`
- `users.limit_invalid_input`
- `users.limit_boundary_min`
- `users.limit_boundary_max`

For an enum parameter `status` with values `["active", "inactive"]`:
- `user.status_valid_active`
- `user.status_valid_inactive`
- `user.status_invalid_input`

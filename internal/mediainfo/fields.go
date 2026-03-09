package mediainfo

func appendFieldUnique(fields []Field, field Field) []Field {
	for _, existing := range fields {
		if existing.Name == field.Name {
			return fields
		}
	}
	return append(fields, field)
}

func setFieldValue(fields []Field, name, value string) []Field {
	for i := range fields {
		if fields[i].Name == name {
			fields[i].Value = value
			return fields
		}
	}
	return append(fields, Field{Name: name, Value: value})
}

func insertFieldBefore(fields []Field, field Field, before string) []Field {
	for i := range fields {
		if fields[i].Name == field.Name {
			fields[i].Value = field.Value
			return fields
		}
	}
	for i, existing := range fields {
		if existing.Name == before {
			fields = append(fields, Field{})
			copy(fields[i+1:], fields[i:])
			fields[i] = field
			return fields
		}
	}
	return append(fields, field)
}

func insertFieldAfter(fields []Field, field Field, after string) []Field {
	for i := range fields {
		if fields[i].Name == field.Name {
			fields[i].Value = field.Value
			return fields
		}
	}
	for i, existing := range fields {
		if existing.Name == after {
			pos := i + 1
			fields = append(fields, Field{})
			copy(fields[pos+1:], fields[pos:])
			fields[pos] = field
			return fields
		}
	}
	return append(fields, field)
}

func removeField(fields []Field, name string) []Field {
	for i := range fields {
		if fields[i].Name == name {
			return append(fields[:i], fields[i+1:]...)
		}
	}
	return fields
}

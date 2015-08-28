package main

func getUsername(id string) (string, error) {
	if u, ok := ids2usernames[id]; ok {
		return u, nil
	}

	u, err := api.GetUserInfo(id)
	if err == nil {
		ids2usernames[id] = u.Name
		return "", err
	}
	return u.Name, nil
}

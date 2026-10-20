package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	flag.Parse()
	commandLineArgs := flag.Args()
	username, valErr := validate(commandLineArgs)
	if valErr != nil {
		fmt.Println(valErr)
		return
	}
	events, apiErr := api(username)
	if apiErr != nil {
		fmt.Println(apiErr)
		return
	}

	output := parser(events)
	for i, o := range output {
		fmt.Println(o)
		if i == 5 {
			fmt.Println("...")
			return
		}
	}
}

func validate(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("no username provided")
	} else if len(args) > 1 {
		return "", errors.New("too many arguments")
	}

	return args[0], nil
}

func api(username string) ([]Event, error) {

	resp, err := http.Get(urlConstructor(username))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	events := []Event{}
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return events, nil
}

func parser(events []Event) []string {
	var output []string
	for _, event := range events {
		switch event.Type {
		case EventTypePush:
			output = append(output,
				fmt.Sprintf("Pushed commits to %s", event.Repo.Name),
			)
		case EventTypeCreate:
			output = append(output,
				fmt.Sprintf("Created in %s", event.Repo.Name),
			)
		case EventTypeDelete:
			output = append(output,
				fmt.Sprintf("Deleted in %s", event.Repo.Name),
			)
		case EventTypePublic:
			output = append(output,
				fmt.Sprintf("Made %s public", event.Repo.Name),
			)
		case EventTypeRelease:
			output = append(output,
				fmt.Sprintf("Published a release in %s", event.Repo.Name),
			)
		}
	}
	return output
}

func urlConstructor(username string) string {
	const url = "https://api.github.com/users/"
	return url + username + "/events"
}

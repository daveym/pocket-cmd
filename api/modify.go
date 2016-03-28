package api

import "github.com/daveym/lint/pocket"

// Modify against Pocket API. Interface used to allow mock to be passed in.
func Modify(pc pocket.API, action string, itemVal int, args []string) string {

	msg := ""

	if len(pc.GetConsumerKey()) == 0 {
		msg = "Consumer Key is not present in lint.yaml. Please add, using the format 'ConsumerKey: value', without quotes."
		return msg
	}

	switch action {
	case "archive":
		msg = applyAction(pc, action, itemVal, args)
	}

	return msg
}

func applyAction(pc pocket.API, action string, itemVal int, args []string) string {

	msg := ""
	modreq := pocket.ModifyRequest{}
	modreq.ConsumerKey = pc.GetConsumerKey()
	modreq.AccessToken = pc.GetAccessToken()

	modact := &pocket.Action{}
	modact.Action = action
	modact.ItemID = itemVal
	modreq.Actions = append(modreq.Actions, modact)

	modresp := &pocket.ModifyResponse{}

	err := pc.Modify(modreq, modresp)

	if err != nil {
		msg = "Error communicating with Pocket: " + err.Error()
		return msg
	}

	return "Update applied successfully"
}

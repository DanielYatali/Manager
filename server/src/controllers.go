package src

type Employee struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	JobTitle  string `json:"jobTitle"`
	Email     string `json:"email"`
	Dob       string `json:"dob"`
	Age       string `json:"age"`
}

// func EmployeeController(api weos.Container, projection projections.Projection, commandDispatcher model.CommandDispatcher, eventSource model.EventRepository, entityFactory model.EntityFactory) echo.HandlerFunc {
// 	return func(ctxt echo.Context) error {
// 		if baseURL, ok := ctxt.Request().Context().Value("baseURL").(string); ok {
// 			if apiURL, ok := ctxt.Request().Context().Value("apiURL").(string); ok {
// 				ctxt.Logger().Debugf("apiURL: '%s'", apiURL)
// 				payload := context.GetPayload(ctxt.Request().Context())
// 				requestData := make(map[string]interface{})
// 				json.Unmarshal(payload, &requestData)
// 				//call the fhir api to send the information

// 				httpClient, err := api.GetHTTPClient("Default")
// 				if err != nil {

// 				}
// 				EmployeePayload := &Employee{}
// 				json.Unmarshal(payload, EmployeePayload)
// 				patient.Name = append(patient.Name, name)
// 				if phone, ok := requestData["phoneNumber"].(string); ok && phone != "" {
// 					patient.Telecom = append(patient.Telecom, ContactPoint{
// 						System: "phone",
// 						Value:  phone,
// 					})
// 				}

// 				if email, ok := requestData["email"].(string); ok && email != "" {
// 					patient.Telecom = append(patient.Telecom, ContactPoint{
// 						System: "email",
// 						Value:  email,
// 					})
// 				}
// 				EmployeePayload.Participant = append(EmployeePayload.Participant, Participant{
// 					Actor: patient,
// 				})

// 				requestedTime, err := time.Parse("2006-01-02 3:04pm", requestData["EmployeeDate"].(string)+" "+requestData["time"].(string))
// 				if err != nil {
// 					ctxt.Logger().Debugf("'%s',invalid date format '%s'", requestData["EmployeeDate"].(string)+" "+requestData["time"].(string), err)
// 				}
// 				EmployeePayload.Start = model.Time{requestedTime}
// 				EmployeePayload.PatientInstruction = requestData["message"].(string)
// 				EmployeePayload.Status = "proposed"
// 				EmployeeData, _ := json.Marshal(EmployeePayload)
// 				body := bytes.NewReader(EmployeeData)
// 				resp, err := httpClient.Post(apiURL+"/Employee", "application/json", body)
// 				if err != nil {
// 					ctxt.Logger().Errorf("unexpected calling the FHIR api '%s'", err)
// 				}
// 				if resp.StatusCode == http.StatusCreated {
// 					var respPayload map[string]interface{}
// 					payloadRAW, err := io.ReadAll(resp.Body)
// 					if err != nil {
// 						ctxt.Logger().Debugf("unable to read response '%s'", err)
// 					}
// 					err = json.Unmarshal(payloadRAW, &respPayload)
// 					if id, ok := respPayload["id"].(float64); ok {
// 						return ctxt.Redirect(http.StatusSeeOther, baseURL+"/patient?Employee="+strconv.Itoa(int(id)))
// 					} else {
// 						ctxt.Logger().Debugf("expected id to be returned and be int64")
// 						return ctxt.Redirect(http.StatusSeeOther, baseURL+"/patient")
// 					}

// 				} else {
// 					return ctxt.Redirect(http.StatusSeeOther, baseURL+"/Employee.html")
// 				}
// 			} else {
// 				ctxt.Logger().Errorf("API_URL environment variable not configured")
// 				return ctxt.NoContent(http.StatusInternalServerError)
// 			}
// 		} else {
// 			ctxt.Logger().Errorf("BASE_URL environment variable not configured")
// 			return ctxt.NoContent(http.StatusInternalServerError)
// 		}

// 	}
// }

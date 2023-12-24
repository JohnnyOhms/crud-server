package Doc

import (
	"fmt"
	"net/http"
)

func Documentation(w http.ResponseWriter, r *http.Request) {
	doc := `
<style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            padding: 20px;
        }
        h1 {
            color: #333;
        }
        h3 {
            color: #555;
        }
        pre {
            background-color: #f5f5f5;
            padding: 10px;
            border-radius: 5px;
            overflow-x: auto;
        }
    </style>
<body>
    <h1>Documentation:</h1>
    <p>For the following API endpoints, observe the guidelines and specifications provided for seamless operation:</p>

    <h3>---> Signup (POST Request)</h3>
    <p>URL: <a href="#">https://crud-api-s9wj.onrender.com/signup/</a></p>
    <p>Return a data structure along with your secret userId.</p>
       <p>Structure to send in the request body:</p>
    <pre>
{
    "UserName": "string",
    "Email": "string",
    "Password": "string",
}
</pre>
<p>convert your password to string</p>

    <h3>---> Signin (POST Request)</h3>
    <p>URL: <a href="#">https://crud-api-s9wj.onrender.com/signin/</a></p>
    <p>Return a data structure along with your secret userId.</p>
    <p>Structure to send in the request body:</p>
    <pre>
{
    "Email": "string",
    "Password": "string",
}
</pre>
<p>convert your password to string </p>

    <p><strong>Note:</strong> Upon authenticating a user, retrieve the userId. Store it locally and include it in every subsequent request. Avoid displaying the userId on the dashboard as it serves as your secret key for all requests.</p>

    <h3>---> Create Info (POST Request)</h3>
    <p>URL: <a href="#">https://crud-api-s9wj.onrender.com/createinfo/</a></p>
    <p>Structure to send in the request body:</p>
    <pre>{
    "UserId": "string",
    "Name": "string",
    "Email": "string",
    "Phone": int,
    "Address": "string",
    "DateCreated": "string" (Ensure your new date format is converted to string format)
}</pre>

    <h3>---> Get All Info (GET Request)</h3>
    <p>URL: <a href="#">https://crud-api-s9wj.onrender.com/getinfo/{userId}/</a></p>
    <p>Send the userId as a parameter in the GET Request.</p>

    <h3>---> Get Single Info (GET Request)</h3>
    <p>URL: <a href="#">https://crud-api-s9wj.onrender.com/getsingleinfo/{userId}/{id}/</a></p>
    <p>Send the userId and the info's ID as parameters in the GET Request.</p>

    <h3>---> Update Info (PUT Request)</h3>
    <p>URL: <a href="#">https://crud-api-s9wj.onrender.com/editinfo/{userId}/{id}/</a></p>
    <p>Send the userId and the info's ID as parameters in the PUT Request.</p>
    <p>Structure to send in the request body:</p>
    <pre>{
    "any": "your preferred updated body key"
}</pre>

    <h3>---> Delete Single Info (DELETE Request)</h3>
    <p>URL: <a href="#">https://crud-api-s9wj.onrender.com/deletesingleinfo/{userId}/{id}/</a></p>
    <p>Send the userId and the info's ID as parameters in the DELETE Request.</p>

    <h3>---> Delete All Info (DELETE Request)</h3>
    <p>URL: <a href="#">https://crud-api-s9wj.onrender.com/deleteinfo/{userId}/</a></p>
    <p>Send the userId as a parameter in the DELETE Request.</p>
</body>`

	w.Header().Set("Content-Type:", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, doc)
}

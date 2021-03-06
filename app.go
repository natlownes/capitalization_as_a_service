package app

import (
	"fmt"
	"net/http"
	"strings"
)

func CapitalizeHandler(w http.ResponseWriter, r *http.Request) {
	var input string

	if r.Method == "GET" {
		input = r.FormValue("arg")
	}

	if r.Method == "POST" {
		if r.ContentLength == 0 {
			http.Error(w, "HTTP Body required", 400)
			return
		}

		body := make([]byte, r.ContentLength, r.ContentLength*100)
		_, err := r.Body.Read(body)

		if err != nil {
			http.Error(w, "Failed reading body", 400)
			return
		}
		input = string(body)
	}

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, strings.ToUpper(input))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	body := `
  <html>
    <head>
      <title>
        Capitalization Microsevice Documentation
      </title>
      <style type="text/css">
        body {
          margin-left:   15%;
          margin-right:  15%;
          padding-bottom: 20px;
        }

        a {
          color:            black;
          text-decoration:  none;
          outline:          none;
        }

        a:hover {
          color:            FF4500;
          text-decoration:  none;
        }

        hr {
          border: 1px solid black;
        }

        code {
          padding:        10px;
          margin-top:     10px;
          margin-bottom:  10px;
          background:     black;
          border:         2px dotted green;
          color:          white;
          font-weight:    900;
          display:        inline-block;
        }

        blockquote {
          padding:      20px;
          background:   lightgrey;
          font-family:  monospace;
        }

        .function h3 {
          font-size:   2em;
          font-family: monospace;
          font-weight: 900;
        }
      </style>
    </head>
    <body>
      <h1><strong>CAPITALIZATION MICROSERVICE</strong></h1>

      <p>
        Everyone knows that microservices are resilent because they exist and
        words are said about them.  Trying to manage capitalization across
        different architectures, languages, and character encodings in a
        performant manner has been a massive stumbling block for modern
        distributed systems and distributed teams.
      </p>

      <p>
        We are excited to announce a powerful, flexible, heavily tested, and
        completely free capitalization microservice for use in any and all of
        your distributed systems.
      </p>

      <div class='function'>
        <hr />

        <h3>
          GET <a href="/capitalize?arg=h" target="_blank">/capitalize?arg=h</a>
        </h3>

        <p>
          Responds with the capitalized version of the letter you request:
        </p>

        <blockquote>
          H
        </blockquote>
      </div>

      <div class='function'>
        <hr />

        <h3>
          GET <a href="/capitalize?arg=hello%2C%20how%20are%20you%3F" target="_blank">
            /capitalize?arg=hello, how are you?
          </a>
        </h3>

        <p>
          You are not restricted to capitalizing one letter at a time (though you
          may find it more performant to do so in a distributed system) - you can
          also request an entire string to be capitalized.
        </p>

        <blockquote>
          HELLO, HOW ARE YOU?
        </blockquote>
      </div>

      <div class='function'>
        <hr />

        <h3>
          POST /capitalize
        </h3>

        <p>
          This too.  Responds with the capitalized body of the request.

          <h5>
            Example:
          </h5>

          <code>
            cat /usr/share/dict/american-english | curl -X POST --data-binary @- http://strings.microservice.narf.io/capitalize
          </code>

          <h5>
            Response:
          </h5>
          <blockquote>
            DOGS<br />
            FREEDOM<br />
            HOTDOGS<br />
            WEAPONS<br />
          </blockquote>
        </p>
      </div>

    </body>
  </html>
  `
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprint(w, body)
}

func CapitalizationService() {
	http.HandleFunc("/capitalize", CapitalizeHandler)
	http.HandleFunc("/", HomeHandler)
}

import "./Download.css";
import { useEffect, useRef, useState } from "react";
import * as App from "../../../wailsjs/go/main/App";
import Searching from "./download/Searching";
import ErrorResponse from "./download/ErrorResponse";
import Found from "./download/Found";

function Download() {
  const inputRef = useRef("");
  const [isSearching, setIsSearching] = useState(false);
  const [isErrorRespounce, setIsErrorRespounce] = useState(false);
  const [responseMessage, setResponseMessage] = useState("");
  const [responseCode, setResponseCode] = useState(0);
  const [responseURL, setresponseURL] = useState("");
  const [responseFaviconURL, setresponseFaviconURL] = useState("");
  const [isFound, setIsFound] = useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    inputRef.current = e.target.value;
  };
  const handleBlur = (e: React.FocusEvent<HTMLInputElement, Element>) => {
    e.target.placeholder = "Enter URL";
  };
  const handleFocus = (e: React.FocusEvent<HTMLInputElement, Element>) => {
    e.target.placeholder = "";
  };
  const handleClick = () => {
    setIsErrorRespounce(false);
    setIsFound(false);
    setIsSearching(true);
    App.FindURL(inputRef.current).then((data) => {
      // set returned values
      setResponseCode(data.code);
      setResponseMessage(data.message);
      setresponseURL(data.url);
      setresponseFaviconURL(data.favicon_url);
      // check for error
      if (data.code > 0) {
        setIsErrorRespounce(true);
      } else if (data.code == 0) {
        setIsFound(true);
      }
      setIsSearching(false);
    });
  };

  return (
    <>
      <div id="input">
        <input
          id="inputArea"
          placeholder="Enter URL"
          onFocus={(e) => {
            handleFocus(e);
          }}
          onBlur={(e) => {
            handleBlur(e);
          }}
          onChange={(e) => handleChange(e)}
        ></input>
        <div id="inputButton" onClick={() => handleClick()}>
          Find
        </div>
      </div>
      <div id="nextAction">
        {isSearching && <Searching />}
        {isErrorRespounce && (
          <ErrorResponse
            message={responseMessage}
            code={responseCode}
            url={responseURL}
            favicon_url={responseFaviconURL}
          />
        )}
        {isFound && (
          <Found
            message={responseMessage}
            code={responseCode}
            url={responseURL}
            favicon_url={responseFaviconURL}
          />
        )}
      </div>
    </>
  );
}

export default Download;

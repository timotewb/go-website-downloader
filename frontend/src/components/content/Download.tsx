import "./Download.css";
import { useEffect, useRef, useState } from "react";
import * as App from "../../../wailsjs/go/main/App";
import Searching from "./download/Searching";
import ErrorResponse from "./download/ErrorResponse";

function Download() {
  const inputRef = useRef("");
  const [isSearching, setIsSearching] = useState(false);
  const [isErrorResounce, setIsErrorResounce] = useState(false);
  const [responseMessage, setResponseMessage] = useState("");
  const [responseCode, setResponseCode] = useState(0);

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
    setIsErrorResounce(false);
    setIsSearching(true);
    App.FindURL(inputRef.current).then((data) => {
      console.log(data);
      // set returned values
      setResponseCode(data.code);
      setResponseMessage(data.message);
      console.log(responseCode, responseMessage);
      // check for error
      if ((data.code = 1)) {
        setIsErrorResounce(true);
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
        {isErrorResounce && (
          <ErrorResponse message={responseMessage} code={responseCode} />
        )}
      </div>
    </>
  );
}

export default Download;

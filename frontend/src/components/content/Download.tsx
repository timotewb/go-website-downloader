import "./Download.css";
import { useEffect, useRef, useState } from "react";
import * as App from "../../../wailsjs/go/main/App";
import Searching from "./download/Searching";

function Download() {
  const inputRef = useRef("");
  const [isSearching, setIsSearching] = useState(false);

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
    setIsSearching(true);
    App.FindURL(inputRef.current).then(() => {
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
      <div>{isSearching && <Searching />}</div>
    </>
  );
}

export default Download;

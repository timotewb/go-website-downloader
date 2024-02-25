import "./Settings.css";
import * as App from "../../../wailsjs/go/main/App";
import { useEffect, useRef, useState } from "react";

function Settings() {

  const [contentDir, setContentDir] = useState("");

  const inputRef = useRef("");
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    inputRef.current = e.target.value;
  };
  const handleBlur = (e: React.FocusEvent<HTMLInputElement, Element>) => {
    e.target.placeholder = contentDir;
  };
  const handleFocus = (e: React.FocusEvent<HTMLInputElement, Element>) => {
    e.target.placeholder = contentDir;
  };

  App.GetSettings().then((data) => {
    // set returned values
    setContentDir(data.content_dir);
  });

  const handleClick = () => {
    console.log("---- testing ----");
  };
  return <>
        <div id="input">
        <input
          id="inputArea"
          placeholder={contentDir}
          onFocus={(e) => {
            handleFocus(e);
          }}
          onBlur={(e) => {
            handleBlur(e);
          }}
          onChange={(e) => handleChange(e)}
        ></input>
        <div id="inputButton" onClick={() => handleClick()}>
          Update
        </div>
      </div>
  </>;
}

export default Settings;

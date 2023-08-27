import { lib } from "../../../../wailsjs/go/models";
import "./ErrorResponse.css";
import errorSVG from "../../../assets/images/error.svg";
import { useState } from "react";

function ErrorResponse(props: lib.ResponseType) {
  const [hideErrorMessage, setHideErrorMessage] = useState(true);
  // identify which error message to use
  let message: string;
  switch (props.code) {
    case 0:
      message = "System error. Code 0 receieved.";
      break;
    case 1:
      message = "Please enter a valid URL (e.g. https://www.website.com/blog)";
      break;
    default:
      message = "System error. No error code receieved.";
      break;
  }

  const handleClick = () => {
    if (hideErrorMessage) {
      setHideErrorMessage(false);
    } else {
      setHideErrorMessage(true);
    }
  };

  const ErrorMessage = () => {
    return props.message;
  };

  return (
    <>
      <div id="messageArea">
        <div className="icon">
          <img src={errorSVG}></img>
        </div>
        <div className="message">{message}</div>
      </div>
      <div className="viewErrorArea">
        <div id="viewErrorMessageButton" onClick={() => handleClick()}>
          View error message
        </div>
        <div id="errorMessageBox">{hideErrorMessage && <ErrorMessage />}</div>
      </div>
    </>
  );
}

export default ErrorResponse;

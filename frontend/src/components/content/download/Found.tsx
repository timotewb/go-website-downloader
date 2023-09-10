import { lib } from "../../../../wailsjs/go/models";
import "./shared.css";
import faviconSVG from "../../../assets/images/favicon-default.svg";

function Found(props: lib.ResponseType) {
  return (
    <>
      Found <p>{props.url}</p>{" "}
      <p>
        <img src={faviconSVG}></img>
      </p>
      <p>{props.favicon_url}</p>
    </>
  );
}

export default Found;

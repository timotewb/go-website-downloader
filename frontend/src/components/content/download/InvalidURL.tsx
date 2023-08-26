import { lib } from "../../../../wailsjs/go/models";

function InvalidURL(props: lib.ResponseType) {
  return <div>InvalidURL {props.message}</div>;
}

export default InvalidURL;

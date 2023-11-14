import {
  JSXElementConstructor,
  ReactElement,
  ReactNode,
  useEffect,
  useState,
} from "react";
import "./Activity.css";
import * as App from "../../../wailsjs/go/main/App";
import * as Models from "../../../wailsjs/go/models";
import { JSX } from "react/jsx-runtime";

function Activity() {
  const [jobRunning, setJobRunning] = useState(false);

  let cat = {} as Models.models.CheckActivityType;
  const activity = [] as JSX.Element[];

  App.CheckActivity().then((data) => {
    if (data.job_count > 0) {
      cat = data;
      console.log(data.data.length);
      data.data.forEach((r, i) => {
        console.log(r.favicon_url);
        activity.push(
          <p>
            {r.favicon_url}
            <br></br>
            {r.url}
          </p>
        );
      });
      setJobRunning(true);
    } else {
      setJobRunning(false);
    }
  });
  console.log(activity);

  useEffect(() => {
    const interval = setInterval(() => {
      App.CheckActivity().then((data) => {
        if (data.job_count > 0) {
          setJobRunning(true);
        } else {
          setJobRunning(false);
        }
      });
    }, 10000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div>
      <p>Activity</p>
      <p>{jobRunning ? <p>{activity}</p> : <p>Job NOT Running</p>}</p>
    </div>
  );
}

export default Activity;

import { useEffect, useState } from "react";
import "./Activity.css";
import * as App from "../../../wailsjs/go/main/App";

function Activity() {

  const [jobRunning, setJobRunning] = useState(false)

  App.CheckIfJobRunning().then((data) => {if(data.job_count>0){setJobRunning(true)} else{setJobRunning(false)}})

  useEffect(() => {
    const interval = setInterval(() => {App.CheckIfJobRunning().then((data) => {if(data.job_count>0){setJobRunning(true)} else{setJobRunning(false)}})}, 10000);
    return () => clearInterval(interval);
   }, []);


return (
<div>
  <p>Activity</p>
  <p>
{jobRunning ? <p>Job Running</p> : <p>Job NOT Running</p>}
</p>
</div>
)

}

export default Activity;

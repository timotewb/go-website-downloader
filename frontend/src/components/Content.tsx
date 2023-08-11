import "./Content.css";
import { useContext } from "react";
import { PageContext, PageContextType } from "../App";

function Content() {
  const pageContext: PageContextType = useContext(PageContext);

  let test: string;
  if (pageContext.download.downloadState) {
    test = "download";
  } else if (pageContext.gallery.galleryState) {
    test = "gallery";
  } else if (pageContext.activity.activityState) {
    test = "activity";
  } else if (pageContext.settings.settingsyState) {
    test = "settings";
  } else {
    test = "download";
  }

  return <div id="content">hello - {test}</div>;
}
export default Content;

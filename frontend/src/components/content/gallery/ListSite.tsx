
import { PageContext, PageContextType } from "../../../App";
import { useContext } from "react";

interface ListSiteProps {
  site_name: string;
}
function ListSite(props: ListSiteProps) {
  const pageContext: PageContextType = useContext(PageContext);

  const handleClick = () => {
    pageContext.gallery.setShowSiteList(false);
  };

  return(
    <div id="gallery">
      <div>{props.site_name}</div>
      <div onClick={handleClick}>back</div>
    </div>
  );
}

export default ListSite;

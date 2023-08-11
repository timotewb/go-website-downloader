import "./MenuBar.css";
import downloadSVG from "../assets/images/download.svg";
import gallerySVG from "../assets/images/gallery.svg";
import activitySVG from "../assets/images/activity.svg";

const setColorHandler = (e: React.MouseEvent<HTMLDivElement, MouseEvent>) => {
  e.currentTarget.classList.toggle("button");
};

function MenuBar() {
  return (
    <div id="menuBar">
      <div
        id="downloadButton"
        className="button"
        onClick={(e) => {
          console.log("Download");
          e.preventDefault();
          setColorHandler(e);
          // e.target.style.backgroundColor = "black";
          console.log(e.target);
        }}
      >
        <img src={downloadSVG}></img> Download
      </div>
      <div
        id="galleryButton"
        className="button"
        onClick={() => {
          console.log("Gallery");
        }}
      >
        <img src={gallerySVG}></img> Gallery
      </div>
      <div
        id="activityButton"
        className="button"
        onClick={() => {
          console.log("Activity");
        }}
      >
        <img src={activitySVG}></img> Activity
      </div>
    </div>
  );
}
export default MenuBar;
//https://stackoverflow.com/questions/38980051/reactjs-adding-active-class-to-button

import React from "react";
import "./HorizontalDivider.css"
// we could also import a library like classNames to make 
// this more flexible

// we could use different spacing types for a design system type approach
// enum SpacingTypes {
  // None = "none",
  // Tight = "tight",
  // Relaxed = "relaxed"
// }

// type Props = {
  /** Spacing on the top and bottom of the divider */
  // spacing?: string;
// }

const HorizontalDivider = () => {
  return <hr className="HorizontalDivider" />
}

export default HorizontalDivider;
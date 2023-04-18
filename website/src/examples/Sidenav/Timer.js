import { format } from "date-fns";
import PropTypes from "prop-types";
import { useEffect, useState } from "react";
import MDTypography from "../../components/MDTypography";

export default function Timer({ textColor, startTime }) {
  const [time, setTime] = useState(startTime);

  useEffect(() => {
    setInterval(() => {
      setTime(new Date());
    }, 1000);
  }, [time]);

  return (
    <MDTypography
      color={textColor}
      display="block"
      variant="caption"
      fontWeight="bold"
      textTransform="uppercase"
      pl={3}
      mt={2}
      mb={1}
      ml={1}
    >
      Врема: {format(time, "HH:mm:ss")}
    </MDTypography>
  );
}

Timer.defaultProps = {
  textColor: "dark",
};

Timer.propTypes = {
  startTime: PropTypes.instanceOf(Date).isRequired,
  textColor: PropTypes.oneOf([
    "inherit",
    "primary",
    "secondary",
    "info",
    "success",
    "warning",
    "error",
    "light",
    "dark",
    "text",
    "white",
  ]),
};

import { addSeconds, format } from "date-fns";
import PropTypes from "prop-types";
import { useEffect, useState } from "react";
import MDTypography from "../../components/MDTypography";
import { useMaterialUIController } from "../../context";

export default function Timer({ textColor }) {
  const [controller] = useMaterialUIController();
  const { date } = controller;
  const [time, setTime] = useState(date);

  useEffect(() => {
    setTime(date);
  }, [date]);
  useEffect(() => {
    const interval = setInterval(() => {
      const currentTime = addSeconds(time, 1);
      setTime(currentTime);
    }, 1000);
    return () => clearInterval(interval);
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

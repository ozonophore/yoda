import { LocalizationProvider, TimePicker } from "@mui/x-date-pickers";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { useState } from "react";
import MDInput from "../../../components/MDInput";

function TimePickers() {
  const [values, setValues] = useState([]);

  return (
    <LocalizationProvider dateAdapter={AdapterDateFns}>
      <TimePicker
        helperText="Helper text"
        onChange={(item) => {
          console.log(`#${item}`);
          setValues(item);
        }}
        onAccept={(v) => {
          setValues([...values, v.format("HH:mm")]);
        }}
        inputFormat="HH:mm"
        renderInput={(props) => (
          <MDInput type="text" helperText="Helper text" {...props} value="wwer" />
        )}
      />
    </LocalizationProvider>
  );
}

export default TimePickers;

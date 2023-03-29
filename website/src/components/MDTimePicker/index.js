import { forwardRef, useState } from "react";
import TextField from "@mui/material/TextField";
import PopupState, { bindPopover, bindTrigger } from "material-ui-popup-state";
import { LocalizationProvider, StaticTimePicker } from "@mui/x-date-pickers";
import { AdapterMoment } from "@mui/x-date-pickers/AdapterMoment";
import PropTypes from "prop-types";
import { MuiChipsInput } from "mui-chips-input";
import { Popover } from "@mui/material";

const MDTimePicker = forwardRef(({ language, ...rest }, ref) => {
  const [values, setValues] = useState([]);

  return (
    <PopupState variant="popover" {...rest} ref={ref}>
      {(popupState) => (
        <div>
          <MuiChipsInput
            autoComplete="off"
            label="Label"
            title="title"
            style={{
              width: "300px",
            }}
            value={values}
            hideClearAll
            onChange={(item) => {
              console.log(item);
              setValues(item);
            }}
            onDeleteChip={(item) => {
              const newSelectedItem = [...values];
              newSelectedItem.splice(newSelectedItem.indexOf(item), 1);
              setValues(newSelectedItem);
            }}
            {...bindTrigger(popupState)}
          />
          <Popover
            {...bindPopover(popupState)}
            anchorOrigin={{
              vertical: "bottom",
              horizontal: "center",
            }}
            elevation={20}
            transformOrigin={{
              vertical: "top",
              horizontal: "center",
            }}
            PaperProps={{
              style: {
                backgroundColor: "white",
              },
            }}
          >
            <LocalizationProvider dateAdapter={AdapterMoment} adapterLocale={language}>
              <StaticTimePicker
                {...bindPopover(popupState)}
                ampmInClock={false}
                ampm={false}
                displayStaticWrapperAs="desktop"
                onChange={(newValue) => {
                  console.log(newValue);
                }}
                onAccept={(v) => {
                  setValues([...values, v.format("HH:mm")]);
                }}
                renderInput={(params) => <TextField {...params} />}
              />
            </LocalizationProvider>
          </Popover>
        </div>
      )}
    </PopupState>
  );
});

// Setting default values for the props of MDTypography
MDTimePicker.defaultProps = {
  language: "ru",
};

MDTimePicker.propTypes = {
  language: PropTypes.string,
};

export default MDTimePicker;

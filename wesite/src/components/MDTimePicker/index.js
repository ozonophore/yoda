import { forwardRef, useState } from "react";
import TextField from "@mui/material/TextField";
import { Chip, InputAdornment, Popover } from "@mui/material";
import IconButton from "@mui/material/IconButton";
import PopupState, { bindPopover, bindTrigger } from "material-ui-popup-state";
import Icon from "@mui/material/Icon";
import { LocalizationProvider, StaticTimePicker } from "@mui/x-date-pickers";
import { AdapterMoment } from "@mui/x-date-pickers/AdapterMoment";
import PropTypes from "prop-types";
import { MuiChipsInput } from "mui-chips-input";

const MDTimePicker = forwardRef(({ language, ...rest }, ref) => {
  const [values, setValues] = useState([]);

  const handleDelete = (item) => () => {
    const newSelectedItem = [...values];
    newSelectedItem.splice(newSelectedItem.indexOf(item), 1);
    setValues(newSelectedItem);
  };
  return (
    <PopupState variant="popover" {...rest} ref={ref}>
      {(popupState) => (
        <div>
          <MuiChipsInput
            label="Label"
            title="title"
            helperText="Helper"
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
          <TextField
            style={{ width: "100px" }}
            InputProps={{
              style: {
                width: "100px",
              },
              readOnly: true,
              startAdornment: values.map((item) => (
                <Chip key={item} tabIndex={-1} label={item} onDelete={handleDelete(item)} />
              )),
              endAdornment: (
                <InputAdornment position="end">
                  <IconButton edge="end" color="primary" {...bindTrigger(popupState)}>
                    <Icon>more_time</Icon>
                  </IconButton>
                </InputAdornment>
              ),
            }}
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

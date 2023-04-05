import { Visibility, VisibilityOff } from "@mui/icons-material";
import { InputAdornment, Stack } from "@mui/material";
import Avatar from "@mui/material/Avatar";
import Card from "@mui/material/Card";
import Grid from "@mui/material/Grid";
import IconButton from "@mui/material/IconButton";
import PropTypes from "prop-types";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import MDBox from "../../../../components/MDBox";
import MDButton from "../../../../components/MDButton";
import MDInput from "../../../../components/MDInput";
import MDTypography from "../../../../components/MDTypography";

const initialRoom = {
  code: "",
  name: "",
  ozon: {
    clientId: "",
    apiKey: "",
  },
  wb: {
    authToken: "",
  },
};

function RoomCard({ room, onCancel, onSubmit }) {
  const [newRoom, setNewRoom] = useState(room || initialRoom);
  const [showWbPassword, setShowWbPassword] = useState(false);
  const [showOzonPassword, setShowOzonPassword] = useState(false);
  const [disabled, setDisabled] = useState(true);
  useEffect(() => {
    setNewRoom(room || initialRoom);
    setShowOzonPassword(false);
    setShowWbPassword(false);
  }, []);
  const [t] = useTranslation();
  const roomNameTitle = `${t("menu.room")}${room.code ? ` - ${room.code}` : ""}`;

  const validateModel = (r) => {
    const { code, name } = r;
    const { clientId, apiKey } = r.ozon;
    const { authToken } = r.wb;
    setDisabled(
      code?.length <= 0 ||
        name?.length <= 0 ||
        clientId?.length <= 0 ||
        apiKey?.length <= 0 ||
        authToken?.length <= 0
    );
  };
  const handleClickWbShowPassword = (e) => {
    setShowWbPassword(!showWbPassword);
    e.preventDefault();
  };
  const handleMouseDownPassword = (event) => {
    event.preventDefault();
  };

  const handleClickOzonShowPassword = (e) => {
    setShowOzonPassword(!showOzonPassword);
    e.preventDefault();
  };
  const handleOnSubmit = () => {
    onSubmit(newRoom);
  };
  const handleOnEditRoom = (r) => {
    setNewRoom(r);
    validateModel(r);
  };

  const { code, name } = newRoom;
  const { clientId, apiKey } = newRoom.ozon;
  const { authToken } = newRoom.wb;

  return (
    <Card>
      <MDBox
        variant="gradient"
        bgColor="info"
        borderRadius="lg"
        coloredShadow="info"
        mx={2}
        mt={-2}
        p={0}
        mb={2}
        textAlign="center"
      >
        <MDTypography variant="h4" fontWeight="medium" color="white" mt={1}>
          {roomNameTitle}
        </MDTypography>
      </MDBox>
      <MDBox pt={1} pb={1} px={3}>
        <MDBox component="form" role="form">
          <MDBox mb={2}>
            <MDInput
              type="text"
              autoComplete="off"
              label={t("room.input.code.title")}
              inputProps={{ maxLength: 20 }}
              error={code?.length <= 0}
              fullWidth
              onChange={(e) => {
                handleOnEditRoom({ ...newRoom, code: e.target.value });
              }}
              value={code}
            />
          </MDBox>
          <MDBox mb={2}>
            <MDInput
              type="text"
              autoComplete="off"
              label={t("room.input.name.title")}
              error={room.name?.length <= 0}
              inputProps={{ maxLength: 100 }}
              fullWidth
              onChange={(e) => {
                handleOnEditRoom({ ...newRoom, name: e.target.value });
              }}
              value={name}
            />
          </MDBox>
          <MDBox pt={1} px={0}>
            <Grid container spacing={2}>
              <Grid xs={12} md={6} lg={6} item>
                <MDBox>
                  <MDTypography
                    variant="button"
                    fontWeight="medium"
                    color="text"
                    display="flex"
                    alignItems="center"
                  >
                    <Avatar src="/images/ozon.svg">OZ</Avatar>
                    &nbsp;&nbsp;{t("room.settings.connection")}
                  </MDTypography>
                  <MDBox mt={2} mb={2}>
                    <MDInput
                      id="ozon-client-id"
                      type="text"
                      autoComplete="off"
                      label="Client-Id"
                      error={clientId?.length <= 0}
                      value={clientId}
                      onChange={(e) =>
                        handleOnEditRoom({
                          ...newRoom,
                          ozon: { ...newRoom.ozon, clientId: e.target.value },
                        })
                      }
                      fullWidth
                    />
                  </MDBox>
                  <MDBox mb={2}>
                    <MDInput
                      id="ozon-api-key"
                      type={showOzonPassword ? "text" : "password"}
                      autoComplete="off"
                      label="Api-Key"
                      fullWidth
                      error={apiKey?.length <= 0}
                      value={apiKey}
                      onChange={(e) =>
                        handleOnEditRoom({
                          ...newRoom,
                          ozon: { ...newRoom.ozon, apiKey: e.target.value },
                        })
                      }
                      InputProps={{
                        autoComplete: "new-password",
                        endAdornment: (
                          <InputAdornment position="end">
                            <IconButton
                              aria-label="toggle password visibility"
                              onClick={handleClickOzonShowPassword}
                              onMouseDown={handleMouseDownPassword}
                              edge="end"
                            >
                              {showOzonPassword ? <VisibilityOff /> : <Visibility />}
                            </IconButton>
                          </InputAdornment>
                        ),
                      }}
                    />
                  </MDBox>
                </MDBox>
              </Grid>
              <Grid xs={12} md={6} lg={6} item>
                <MDBox>
                  <MDTypography
                    variant="button"
                    fontWeight="medium"
                    color="text"
                    display="flex"
                    alignItems="center"
                  >
                    <Avatar src="/images/wb.png">WB</Avatar>
                    &nbsp;&nbsp;{t("room.settings.connection")}
                  </MDTypography>
                  <MDBox mt={2} mb={2} ml={0}>
                    <MDInput
                      id="wb-auth"
                      type={showWbPassword ? "text" : "password"}
                      autoComplete="off"
                      label="Authorization"
                      fullWidth
                      error={authToken?.length <= 0}
                      value={authToken}
                      onChange={(e) =>
                        handleOnEditRoom({
                          ...newRoom,
                          wb: { ...newRoom.wb, authToken: e.target.value },
                        })
                      }
                      InputProps={{
                        autoComplete: "new-password",
                        endAdornment: (
                          <InputAdornment position="end">
                            <IconButton
                              aria-label="toggle password visibility"
                              onClick={handleClickWbShowPassword}
                              onMouseDown={handleMouseDownPassword}
                              edge="end"
                            >
                              {showWbPassword ? <VisibilityOff /> : <Visibility />}
                            </IconButton>
                          </InputAdornment>
                        ),
                      }}
                    />
                  </MDBox>
                </MDBox>
              </Grid>
            </Grid>
          </MDBox>
        </MDBox>
      </MDBox>
      <MDBox pt={0} pb={3} px={3}>
        <Stack direction="row" justifyContent="end" spacing={2}>
          <MDButton variant="outlined" color="info" onClick={onCancel}>
            {t("cancel")}
          </MDButton>
          <MDButton variant="gradient" disabled={disabled} color="info" onClick={handleOnSubmit}>
            {t("submit")}
          </MDButton>
        </Stack>
      </MDBox>
    </Card>
  );
}

RoomCard.defaultProps = {
  room: initialRoom,
};

RoomCard.propTypes = {
  onCancel: PropTypes.func.isRequired,
  onSubmit: PropTypes.func.isRequired,
  room: PropTypes.shape({
    code: PropTypes.string,
    name: PropTypes.string,
    ozon: PropTypes.shape({
      clientId: PropTypes.string,
      apiKey: PropTypes.string,
    }).isRequired,
    wb: PropTypes.shape({
      authToken: PropTypes.string,
    }).isRequired,
  }),
};

export default RoomCard;

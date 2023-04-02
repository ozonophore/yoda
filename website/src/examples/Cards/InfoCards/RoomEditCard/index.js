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

function RoomEditCard({ code, name, ozon, wb, onCancel }) {
  const [roomCode, setRoomCode] = useState("");
  const [roomName, setRoomName] = useState("");
  const [ozonClientId, setOzonClientId] = useState("");
  const [ozonApiKey, setOzonApiKey] = useState("");
  const [wbAuth, setWbAuth] = useState("");
  const [showWbPassword, setShowWbPassword] = useState(false);
  const [showOzonPassword, setShowOzonPassword] = useState(false);
  useEffect(() => {
    setRoomCode(code);
    setRoomName(name);
    setOzonClientId(ozon.clientId);
    setOzonApiKey(ozon.apiKey);
    setWbAuth(wb.authorization);
    setShowOzonPassword(false);
    setShowWbPassword(false);
  }, [code, name, ozon, wb]);
  const [t] = useTranslation();
  const roomNameTitle = `${t("menu.room")}${roomCode ? ` - ${roomCode}` : ""}`;
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
  const handleOzonClientIdChange = (e) => {
    setOzonClientId(e.target.value);
  };
  const handleOzonApiKeyChange = (e) => {
    setOzonApiKey(e.target.value);
  };
  const handleWbAuthChange = (e) => {
    setWbAuth(e.target.value);
  };

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
              error={roomCode?.length <= 0}
              fullWidth
              onChange={(e) => {
                setRoomCode(e.target.value);
              }}
              value={roomCode}
            />
          </MDBox>
          <MDBox mb={2}>
            <MDInput
              type="text"
              autoComplete="off"
              label={t("room.input.name.title")}
              error={roomName?.length <= 0}
              inputProps={{ maxLength: 100 }}
              fullWidth
              onChange={(e) => {
                setRoomName(e.target.value);
              }}
              value={roomName}
            />
          </MDBox>
          <MDBox pt={1} px={0}>
            <Grid container spacing={0}>
              <Grid xs={12} md={6} lg={6} item>
                <MDBox>
                  <MDBox pb={3}>
                    <Avatar src="/images/ozon.svg">OZ</Avatar>
                  </MDBox>
                  <MDBox mb={2}>
                    <MDInput
                      type="text"
                      autoComplete="off"
                      label="Client-Id"
                      value={ozonClientId}
                      onChange={handleOzonClientIdChange}
                      fullWidth
                    />
                  </MDBox>
                  <MDBox mb={2}>
                    <MDInput
                      type={showOzonPassword ? "text" : "password"}
                      autoComplete="off"
                      label="Api-Key"
                      fullWidth
                      value={ozonApiKey}
                      onChange={handleOzonApiKeyChange}
                      InputProps={{
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
                  <MDBox pb={3} ml={2}>
                    <Avatar src="/images/wb.png">WB</Avatar>
                  </MDBox>
                  <MDBox mb={2} ml={2}>
                    <MDInput
                      type={showWbPassword ? "text" : "password"}
                      autoComplete="off"
                      label="Authorization"
                      fullWidth
                      value={wbAuth}
                      onChange={handleWbAuthChange}
                      InputProps={{
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
          <MDButton variant="gradient" color="info">
            {t("submit")}
          </MDButton>
        </Stack>
      </MDBox>
    </Card>
  );
}

RoomEditCard.propTypes = {
  code: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  onCancel: PropTypes.func.isRequired,
  ozon: PropTypes.shape({
    clientId: PropTypes.string.isRequired,
    apiKey: PropTypes.string.isRequired,
  }).isRequired,
  wb: PropTypes.shape({
    authorization: PropTypes.string.isRequired,
  }).isRequired,
};

export default RoomEditCard;

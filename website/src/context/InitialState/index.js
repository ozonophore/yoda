const InitialState = {
  miniSidenav: false,
  transparentSidenav: false,
  whiteSidenav: false,
  sidenavColor: "info",
  transparentNavbar: true,
  fixedNavbar: true,
  openConfigurator: false,
  direction: "ltr",
  layout: "dashboard",
  darkMode: false,
  language: "ru",
  loading: false,
  room: {
    addToggle: false,
    gridToggle: false,
  },
  rooms: [],
  jobs: [
    {
      id: 1,
      name: "Загрузка данных",
      description: "Загрузка данных из внешних источников",
      loader: {
        weekDays: [],
        atTimes: [],
        interval: null,
        maxRuns: null,
        nextRun: null,
        lastRun: null,
      },
    },
  ],
  error: {
    message: "",
  },
  date: new Date(),
  nextRun: null,
};

export default InitialState;

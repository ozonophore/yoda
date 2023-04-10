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
  rooms: [],
  jobs: {
    loader: {
      weekDays: [],
      atTimes: [],
      nextRun: new Date(),
      lastRun: new Date(),
    },
    addLoader: {
      interval: 0,
      maxRuns: 0,
    },
    calcAggregates: {
      weekDays: [],
      atTimes: [],
      nextRun: null,
      lastRun: null,
    },
  },
  error: {
    message: "",
  },
};

export default InitialState;

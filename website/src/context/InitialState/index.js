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
      id: 0,
      name: "Загрузка данных",
      description: "Загрузка данных из внешних источников",
      loader: {
        weekDays: ["monday", "tuesday", "wednesday", "thursday", "friday"],
        atTimes: ["00:00", "01:00", "02:00", "03:00", "04:00", "05:00"],
        interval: null,
        maxRuns: null,
        nextRun: new Date(),
        lastRun: new Date(),
      },
    },
    {
      id: 1,
      name: "Обнвление заданий",
      description: "Обновление заданий",
      loader: {
        weekDays: null,
        atTimes: null,
        interval: 1000,
        maxRuns: null,
        nextRun: new Date(),
        lastRun: new Date(),
      },
    },
    {
      id: 2,
      name: "Дозагрузка данных",
      description: "Дозагрузка данных из внешних источников",
      loader: {
        weekDays: null,
        atTimes: null,
        interval: 1000,
        maxRuns: 10,
        nextRun: new Date(),
        lastRun: new Date(),
      },
    },
  ],
  error: {
    message: "",
  },
};

export default InitialState;

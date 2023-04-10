import { isFunction } from "lodash";
import { useCallback, useReducer } from "react";

const useThunkReducer = (reducer, initialState) => {
  const [state, dispatch] = useReducer(reducer, initialState);

  // As long as the dispatch is the same, the enhancedDispatch is the same one.
  const enhancedDispatch = useCallback(
    (action) => {
      if (isFunction(action)) {
        action(dispatch);
      } else {
        dispatch(action);
      }
    },
    [dispatch]
  );

  return [state, enhancedDispatch];
};

export default useThunkReducer;

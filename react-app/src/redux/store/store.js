import { configureStore } from '@reduxjs/toolkit';
import accountReducer from '../feature/accountSlice'


export default configureStore({
    reducer: {
        // ここの名前をsliceの末尾にあるreducerの名前と絶対揃える
        accounts: accountReducer,
        //   keywords: keywordsReducer,
    },
  });


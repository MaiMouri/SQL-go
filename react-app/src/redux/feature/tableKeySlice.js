import { createSlice } from '@reduxjs/toolkit';
// import { getAccounts } from "../api";


//createSliceを使ったら reducer を作成するだけで自動的に action type も定義してくれるし action creator も生成。
export const tableKeySlice = createSlice({
  //name -> action type の prefix
  name: "tableKeys",
  initialState: { loadingNow: false, error: null, tableKeys: [] },
  reducers: {
    //addUser -> action type
    // addAccount: (state, action) => {
    //   state.accounts = [action.payload, ...state.accounts]
    // },
    // updateAccount: (state, action) => {
    //   const { AccountId, Word, ImageUrl, Description } = action.payload
    //   const existingAccount = state.accounts.find((keyword) => keyword.AccountId === AccountId);
    //   console.log("upadtaAccount in slice" + existingAccount);
    //   console.log(AccountId, Word, ImageUrl, Description);
    //   // state.accounts = state.accounts.map((item) => item.AccountId === action.payload.AccountId);
    //   if (existingAccount) {
    //     existingAccount.Word = Word;
    //     existingAccount.Description = Description;
    //     existingAccount.ImageUrl = ImageUrl;
    //   }
    // },
    // deleteAccount: (state, action) => {
    //   state.accounts = state.accounts.filter((item) => item.AccountId !== action.payload.AccountId);
    // },
    getTableKey: (state, action) => {
      state.accounts = state.accounts.filter((item) => item.AccountId !== action.payload.AccountId);
    },
     // 通信を開始した時に呼ぶ関数
    fetchStart(state, action) {
      state.loadingNow = true;
      state.error = null;
    },
    // 通信が失敗した時に呼ぶ関数
    fetchFailure(state, action) {
      state.loading = false;
      state.error = action.payload;
    },
    // 通信が成功した時に呼ぶ関数
    fetchSuccess(state, action) {
      state.loadingNow = false;
      state.error = null;
      state.accounts = action.payload;
    },
  },
});

// Actions
export const { fetchStart, fetchFailure, fetchSuccess, getTableKey } = accountSlice.actions;


// 外部からはこの関数を呼んでもらう
export const fetchItems = () => async dispatch => {
  try {
    dispatch(fetchStart());
    dispatch(fetchSuccess(await getTableKey()));
  } catch (error) {
    dispatch(fetchFailure(error.stack));
  }
};

// Selectors keywords-storeのreducer nameと揃える
export const selectAccount = ({ tableKeys }) => tableKeys;

// Reducer(must be default export)
export default tableKeysSlice.reducer;
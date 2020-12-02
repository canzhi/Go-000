1. 如果DAO层，遇到sql.ErrNoRows 应该抛出错误给上层。
if err != nil {
  if _, ok := err.(sql.ErrNoRows); ok {
    return errors.Wrap(err, "DAO: NO ROWS FROM DB")
  } else {
    return errors.Wrap(err, "DAO: not sql.ErrNoRows error")
  }
}

package repository

var (
	queryCreateUser = `
		INSERT INTO Users (ID, FullName, Email, HashedPassword, ProfilePicture)
		VALUE (?, ?, ?, ?, ?);
	`
	queryGetUserByEmail = `
		SELECT ID FROM Users 
		WHERE Email = ? LIMIT 1;
	`
	queryUpdatePassword = `
		UPDATE Users 
		SET HashedPassword = ?
		WHERE ID = ?;
	`
	queryCreateAttempt = `
		INSERT INTO Users (ID, UserID, Token, ValidUntil) 
		VALUE (?, ?, ?, ?);
	`
	queryDeleteOldAttempt = `
		DELETE FROM ResetAttempts
		WHERE Succeed = FALSE AND UserID = ?;
	`
	queryGetAttempt = `
		SELECT ValidUntil FROM ResetAttempts
		WHERE UserID = ? AND Token = ?;
	`
	queryUpdateAttemptStatus = `
		UPDATE ResetAttempts
		SET Succeed = TRUE
		WHERE ID = ?;
	`
)

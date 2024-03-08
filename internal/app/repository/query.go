package repository

var (
	queryCreateUser = `
		INSERT INTO Users (ID, FullName, Email, HashedPassword, ProfilePicture)
		VALUE (:ID, :FullName, :Email, :HashedPassword, :ProfilePicture);
	`
	queryGetUserByParam = `
		SELECT ID, FullName, Email, HashedPassword, ProfilePicture
		FROM Users 
		WHERE %s = ? LIMIT 1;
	`
	queryUpdatePassword = `
		UPDATE Users 
		SET HashedPassword = :Hashed
		WHERE ID = :ID;
	`
	queryCreateAttempt = `
		INSERT INTO Users (ID, UserID, Token, ValidUntil) 
		VALUE (:ID, :UserID, :Token, :ValidUntil);
	`
	queryDeleteOldAttempt = `
		DELETE FROM ResetAttempts
		WHERE Succeed = FALSE AND UserID = ?;
	`
	queryGetAttemptExpiration = `
		SELECT Expiration FROM ResetAttempts
		WHERE UserID = ? AND Token = ?;
	`
	queryUpdateAttemptStatus = `
		UPDATE ResetAttempts
		SET Succeed = TRUE
		WHERE ID = ?;
	`
)

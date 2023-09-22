package glog

import "os"

var logger = Logger{}

// DEFAULT (0): The log entry has no assigned severity level.
//
// Print internally calls logger.Print with the DEFAULT Severity.
// Arguments are handled in the same way as fmt.Print.
func Print(v ...any) {
	logger.Print(v...)
}

// DEFAULT (0): The log entry has no assigned severity level.
//
// Println internally calls logger.Println with the DEFAULT Severity.
// Arguments are handled in the same way as fmt.Println.
func Println(v ...any) {
	logger.Println(v...)
}

// DEFAULT (0): The log entry has no assigned severity level.
//
// Printf internally calls logger.Printf with the DEFAULT Severity.
// Arguments are handled in the same way as fmt.Printf.
func Printf(format string, v ...any) {
	logger.Printf(format, v...)
}

// DEBUG (100): Debug or trace information.
//
// Debug internally calls logger.Debug with the DEBUG Severity.
// Arguments are handled in the same way as fmt.Print.
func Debug(v ...any) {
	logger.Debug(v...)
}

// DEBUG (100): Debug or trace information.
//
// Debugln internally calls logger.Debugln with the DEBUG Severity.
// Arguments are handled in the same way as fmt.Println.
func Debugln(v ...any) {
	logger.Debugln(v...)
}

// DEBUG (100): Debug or trace information.
//
// Debugf internally calls logger.Debugf with the DEBUG Severity.
// Arguments are handled in the same way as fmt.Printf.
func Debugf(format string, v ...any) {
	logger.Debugf(format, v...)
}

// INFO (200): Routine information, such as ongoing status or performance.
//
// Info internally calls logger.Info with the INFO Severity.
// Arguments are handled in the same way as fmt.Print.
func Info(v ...any) {
	logger.Info(v...)
}

// INFO (200): Routine information, such as ongoing status or performance.
//
// Infoln internally calls logger.Infoln with the INFO Severity.
// Arguments are handled in the same way as fmt.Println.
func Infoln(v ...any) {
	logger.Infoln(v...)
}

// INFO (200): Routine information, such as ongoing status or performance.
//
// Infof internally calls logger.Infof with the INFO Severity.
// Arguments are handled in the same way as fmt.Printf.
func Infof(format string, v ...any) {
	logger.Infof(format, v...)
}

// NOTICE (300): Normal but significant events, such as start up, shut down, or a configuration change.
//
// Notice internally calls logger.Notice with the NOTICE Severity.
// Arguments are handled in the same way as fmt.Print.
func Notice(v ...any) {
	logger.Notice(v...)
}

// NOTICE (300): Normal but significant events, such as start up, shut down, or a configuration change.
//
// Noticeln internally calls logger.Noticeln with the NOTICE Severity.
// Arguments are handled in the same way as fmt.Println.
func Noticeln(v ...any) {
	logger.Noticeln(v...)
}

// NOTICE (300): Normal but significant events, such as start up, shut down, or a configuration change.
//
// Noticef internally calls logger.Noticef with the NOTICE Severity.
// Arguments are handled in the same way as fmt.Printf.
func Noticef(format string, v ...any) {
	logger.Noticef(format, v...)
}

// WARNING (400): Warning events might cause problems.
//
// Warning internally calls logger.Warning with the WARNING Severity.
// Arguments are handled in the same way as fmt.Print.
func Warning(v ...any) {
	logger.Warning(v...)
}

// WARNING (400): Warning events might cause problems.
//
// Warningln internally calls logger.Warningln with the WARNING Severity.
// Arguments are handled in the same way as fmt.Println.
func Warningln(v ...any) {
	logger.Warningln(v...)
}

// WARNING (400): Warning events might cause problems.
//
// Warningf internally calls logger.Warningf with the WARNING Severity.
// Arguments are handled in the same way as fmt.Printf.
func Warningf(format string, v ...any) {
	logger.Warningf(format, v...)
}

// ERROR (500): Error events are likely to cause problems.
//
// Error internally calls logger.Error with the ERROR Severity.
// Arguments are handled in the same way as fmt.Print.
func Error(v ...any) {
	logger.Error(v...)
}

// ERROR (500): Error events are likely to cause problems.
//
// Errorln internally calls logger.Errorln with the ERROR Severity.
// Arguments are handled in the same way as fmt.Println.
func Errorln(v ...any) {
	logger.Errorln(v...)
}

// ERROR (500): Error events are likely to cause problems.
//
// Errorf internally calls logger.Errorf with the ERROR Severity.
// Arguments are handled in the same way as fmt.Printf.
func Errorf(format string, v ...any) {
	logger.Errorf(format, v...)
}

// CRITICAL (600): Critical events cause more severe problems or outages.
//
// Critical internally calls logger.Critical with the CRITICAL Severity.
// Arguments are handled in the same way as fmt.Print.
func Critical(v ...any) {
	logger.Critical(v...)
}

// CRITICAL (600): Critical events cause more severe problems or outages.
//
// Criticalln internally calls logger.Criticalln with the CRITICAL Severity.
// Arguments are handled in the same way as fmt.Println.
func Criticalln(v ...any) {
	logger.Criticalln(v...)
}

// CRITICAL (600): Critical events cause more severe problems or outages.
//
// Criticalf internally calls logger.Criticalf with the CRITICAL Severity.
// Arguments are handled in the same way as fmt.Printf.
func Criticalf(format string, v ...any) {
	logger.Criticalf(format, v...)
}

// ALERT (700): A person must take an action immediately.
//
// Alert internally calls logger.Alert with the ALERT Severity.
// Arguments are handled in the same way as fmt.Print.
func Alert(v ...any) {
	logger.Alert(v...)
}

// ALERT (700): A person must take an action immediately.
//
// Alertln internally calls logger.Alertln with the ALERT Severity.
// Arguments are handled in the same way as fmt.Println.
func Alertln(v ...any) {
	logger.Alertln(v...)
}

// ALERT (700): A person must take an action immediately.
//
// Alertf internally calls logger.Alertf with the ALERT Severity.
// Arguments are handled in the same way as fmt.Printf.
func Alertf(format string, v ...any) {
	logger.Alertf(format, v...)
}

// EMERGENCY (800): One or more systems are unusable.
//
// Emergency internally calls logger.Emergency with the EMERGENCY Severity.
// Arguments are handled in the same way as fmt.Print.
func Emergency(v ...any) {
	logger.Emergency(v...)
}

// EMERGENCY (800): One or more systems are unusable.
//
// Emergencyln internally calls logger.Emergencyln with the EMERGENCY Severity.
// Arguments are handled in the same way as fmt.Println.
func Emergencyln(v ...any) {
	logger.Emergencyln(v...)
}

// EMERGENCY (800): One or more systems are unusable.
//
// Emergencyf internally calls logger.Emergencyf with the EMERGENCY Severity.
// Arguments are handled in the same way as fmt.Printf.
func Emergencyf(format string, v ...any) {
	logger.Emergencyf(format, v...)
}

// EMERGENCY (800): One or more systems are unusable.
//
// Fatal internally calls logger.Fatal with the EMERGENCY Severity followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func Fatal(v ...any) {
	logger.Fatal(v...)
}

// EMERGENCY (800): One or more systems are unusable.
//
// Fatalln internally calls logger.Fatalln with the EMERGENCY Severity followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Println.
func Fatalln(v ...any) {
	logger.Fatalln(v...)
}

// EMERGENCY (800): One or more systems are unusable.
//
// Fatalf internally calls logger.Fatalf with the EMERGENCY Severity followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, v ...any) {
	logger.Fatalf(format, v...)
}

type Logger struct {
	callers     int
	trace       string
	spanID      string
	executionID string
	request     *HttpRequest
}

// DEFAULT (0): The log entry has no assigned severity level.
//
// Print internally calls glog.print with the DEFAULT Severity.
// Arguments are handled in the same way as fmt.Print.
func (l Logger) Print(v ...any) {
	print(DEFAULT, l, v...)
}

// DEFAULT (0): The log entry has no assigned severity level.
//
// Println internally calls glog.println with the DEFAULT Severity.
// Arguments are handled in the same way as fmt.Println.
func (l Logger) Println(v ...any) {
	println(DEFAULT, l, v...)
}

// DEFAULT (0): The log entry has no assigned severity level.
//
// Printf internally calls glog.printf with the DEFAULT Severity.
// Arguments are handled in the same way as fmt.Printf.
func (l Logger) Printf(format string, v ...any) {
	printf(DEFAULT, l, format, v...)
}

// DEBUG (100): Debug or trace information.
//
// Debug internally calls glog.print with the DEBUG Severity.
// Arguments are handled in the same way as fmt.Print.
func (l Logger) Debug(v ...any) {
	print(DEBUG, l, v...)
}

// DEBUG (100): Debug or trace information.
//
// Debugln internally calls glog.println with the DEBUG Severity.
// Arguments are handled in the same way as fmt.Println.
func (l Logger) Debugln(v ...any) {
	println(DEBUG, l, v...)
}

// DEBUG (100): Debug or trace information.
//
// Debugf internally calls glog.printf with the DEBUG Severity.
// Arguments are handled in the same way as fmt.Printf.
func (l Logger) Debugf(format string, v ...any) {
	printf(DEBUG, l, format, v...)
}

// INFO (200): Routine information, such as ongoing status or performance.
//
// Info internally calls glog.print with the INFO Severity.
// Arguments are handled in the same way as fmt.Print.
func (l Logger) Info(v ...any) {
	print(INFO, l, v...)
}

// INFO (200): Routine information, such as ongoing status or performance.
//
// Infoln internally calls glog.println with the INFO Severity.
// Arguments are handled in the same way as fmt.Println.
func (l Logger) Infoln(v ...any) {
	println(INFO, l, v...)
}

// INFO (200): Routine information, such as ongoing status or performance.
//
// Infof internally calls glog.printf with the INFO Severity.
// Arguments are handled in the same way as fmt.Printf.
func (l Logger) Infof(format string, v ...any) {
	printf(INFO, l, format, v...)
}

// NOTICE (300): Normal but significant events, such as start up, shut down, or a configuration change.
//
// Notice internally calls glog.print with the NOTICE Severity.
// Arguments are handled in the same way as fmt.Print.
func (l Logger) Notice(v ...any) {
	print(NOTICE, l, v...)
}

// NOTICE (300): Normal but significant events, such as start up, shut down, or a configuration change.
//
// Noticeln internally calls glog.println with the NOTICE Severity.
// Arguments are handled in the same way as fmt.Println.
func (l Logger) Noticeln(v ...any) {
	println(NOTICE, l, v...)
}

// NOTICE (300): Normal but significant events, such as start up, shut down, or a configuration change.
//
// Noticef internally calls glog.printf with the NOTICE Severity.
// Arguments are handled in the same way as fmt.Printf.
func (l Logger) Noticef(format string, v ...any) {
	printf(NOTICE, l, format, v...)
}

// WARNING (400): Warning events might cause problems.
//
// Warning internally calls glog.print with the WARNING Severity.
// Arguments are handled in the same way as fmt.Print.
func (l Logger) Warning(v ...any) {
	print(WARNING, l, v...)
}

// WARNING (400): Warning events might cause problems.
//
// Warningln internally calls glog.println with the WARNING Severity.
// Arguments are handled in the same way as fmt.Println.
func (l Logger) Warningln(v ...any) {
	println(WARNING, l, v...)
}

// WARNING (400): Warning events might cause problems.
//
// Warningf internally calls glog.printf with the WARNING Severity.
// Arguments are handled in the same way as fmt.Printf.
func (l Logger) Warningf(format string, v ...any) {
	printf(WARNING, l, format, v...)
}

// ERROR (500): Error events are likely to cause problems.
//
// Error internally calls glog.print with the ERROR Severity.
// Arguments are handled in the same way as fmt.Print.
func (l Logger) Error(v ...any) {
	print(ERROR, l, v...)
}

// ERROR (500): Error events are likely to cause problems.
//
// Errorln internally calls glog.println with the ERROR Severity.
// Arguments are handled in the same way as fmt.Println.
func (l Logger) Errorln(v ...any) {
	println(ERROR, l, v...)
}

// ERROR (500): Error events are likely to cause problems.
//
// Errorf internally calls glog.printf with the ERROR Severity.
// Arguments are handled in the same way as fmt.Printf.
func (l Logger) Errorf(format string, v ...any) {
	printf(ERROR, l, format, v...)
}

// CRITICAL (600): Critical events cause more severe problems or outages.
//
// Critical internally calls glog.print with the CRITICAL Severity.
// Arguments are handled in the same way as fmt.Print.
func (l Logger) Critical(v ...any) {
	print(CRITICAL, l, v...)
}

// CRITICAL (600): Critical events cause more severe problems or outages.
//
// Criticalln internally calls glog.println with the CRITICAL Severity.
// Arguments are handled in the same way as fmt.Println.
func (l Logger) Criticalln(v ...any) {
	println(CRITICAL, l, v...)
}

// CRITICAL (600): Critical events cause more severe problems or outages.
//
// Criticalf internally calls glog.printf with the CRITICAL Severity.
// Arguments are handled in the same way as fmt.Printf.
func (l Logger) Criticalf(format string, v ...any) {
	printf(CRITICAL, l, format, v...)
}

// ALERT (700): A person must take an action immediately.
//
// Alert internally calls glog.print with the ALERT Severity.
// Arguments are handled in the same way as fmt.Print.
func (l Logger) Alert(v ...any) {
	print(ALERT, l, v...)
}

// ALERT (700): A person must take an action immediately.
//
// Alertln internally calls glog.println with the ALERT Severity.
// Arguments are handled in the same way as fmt.Println.
func (l Logger) Alertln(v ...any) {
	println(ALERT, l, v...)
}

// ALERT (700): A person must take an action immediately.
//
// Alertf internally calls glog.printf with the ALERT Severity.
// Arguments are handled in the same way as fmt.Printf.
func (l Logger) Alertf(format string, v ...any) {
	printf(ALERT, l, format, v...)
}

// EMERGENCY (800): One or more systems are unusable.
//
// Emergency internally calls glog.print with the EMERGENCY Severity.
// Arguments are handled in the same way as fmt.Print.
func (l Logger) Emergency(v ...any) {
	print(EMERGENCY, l, v...)
}

// EMERGENCY (800): One or more systems are unusable.
//
// Emergencyln internally calls glog.println with the EMERGENCY Severity.
// Arguments are handled in the same way as fmt.Println.
func (l Logger) Emergencyln(v ...any) {
	println(EMERGENCY, l, v...)
}

// EMERGENCY (800): One or more systems are unusable.
//
// Emergencyf internally calls glog.printf with the EMERGENCY Severity.
// Arguments are handled in the same way as fmt.Printf.
func (l Logger) Emergencyf(format string, v ...any) {
	printf(EMERGENCY, l, format, v...)
}

// EMERGENCY (800): One or more systems are unusable.
//
// Fatal internally calls glog.print with the EMERGENCY Severity followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func (l Logger) Fatal(v ...any) {
	print(EMERGENCY, l, v...)
	os.Exit(1)
}

// EMERGENCY (800): One or more systems are unusable.
//
// Fatalln internally calls glog.println with the EMERGENCY Severity followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Println.
func (l Logger) Fatalln(v ...any) {
	println(EMERGENCY, l, v...)
	os.Exit(1)
}

// EMERGENCY (800): One or more systems are unusable.
//
// Fatalf internally calls glog.printf with the EMERGENCY Severity followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func (l Logger) Fatalf(format string, v ...any) {
	printf(EMERGENCY, l, format, v...)
	os.Exit(1)
}

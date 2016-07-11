DESCRIPTION
===========

This tracker listens and parse requests from other tracker, without setting its
own client-side state on user's browser. In this case, other tracker means
tracker type - [bce].

For the coding part, this tracker is simply a listener which waits for other
tracker send information together with the request. Notice this tracker owns an
identifier (`idd\_max`), and this is shared among other trackers in memory.

The important modifications should be made in other trackers. If specified that
type - d tracker is used, then other tracker should serve a page that send
request to type - d tracker.

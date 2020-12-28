<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>tus-js-client demo - File Upload</title>
    <link href="https://netdna.bootstrapcdn.com/twitter-bootstrap/2.3.1/css/bootstrap-combined.min.css" rel="stylesheet">
    <link href="../static/css/demo.css" rel="stylesheet" media="screen">
  </head>
  <body>
    <div class="container">
      <h1>tus-js-client demo - File Upload</h1>

      <p>
        This demo shows the basic functionality of the tus protocol. You can select a file using the controls below and start/pause the upload as you wish.
      </p>

      <p>
        For a prettier demo please go to the
        <a href="http://tus.io/demo.html">tus.io</a> website.
        This demo is just here to aid developers.
      </p>

      <p>
        A demo where a video is recorded from your webcam while being simultaneously uploaded, can be found <a href="./video.html">here</a>.
      </p>

      <div class="alert alert-warining hidden" id="support-alert">
        <b>Warning!</b> Your browser does not seem to support the features necessary to run tus-js-client. The buttons below may work but probably will fail silently.
      </div>

      <br>

      <table>
        <tr>
          <td>
            <label for="endpoint">
              Upload endpoint:
            </label>
          </td>
          <td>
            <input type="text" id="endpoint" name="endpoint" value="http://myproject-beego-test.accelerite-openshift-la-fbc03b92adfe8eb26bb2ca99edfad3f7-0000.che01.containers.appdomain.cloud/files/">
          </td>
        </tr>
        <tr>
          <td>
            <label for="chunksize">
              Chunk size (bytes):
            </label>
          </td>
          <td>
            <input type="number" id="chunksize" name="chunksize">
          </td>
        </tr>
        <tr>
          <td>
            <label for="paralleluploads">
              Parallel upload requests:
            </label>
          </td>
          <td>
            <input type="number" id="paralleluploads" name="paralleluploads" value="1">
          </td>
        </tr>
      </table>

      <br>

      <input type="file">

      <br>
      <br>

      <div class="row">
        <div class="span8">
          <div class="progress progress-striped progress-success">
            <div class="bar" style="width: 0%;"></div>
          </div>
        </div>
        <div class="span4">
          <button class="btn stop" id="toggle-btn">start upload</button>
        </div>
      </div>

      <hr>
      <h3>Uploads</h3>
      <p id="upload-list">
        Succesful uploads will be listed here. Try one!
      </p>

    </div>
  </body>

  <script src="../static/js/tus.js"></script>
  <script src="../static/js/demo.js"></script>
</html>

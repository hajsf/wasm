const SQLpath = '/www/wasm/sql';
const sqlPromise = initSqlJs({
    locateFile: file => `${SQLpath}/sql-wasm.wasm`
  });
const dataPromise = fetch("http://localhost:8080/www/data/database.sqlite").then(res => res.arrayBuffer());
sqlRun();
async function sqlRun(){
  const [SQL, buf] = await Promise.all([sqlPromise, dataPromise])

  const db = new SQL.Database(new Uint8Array(buf));

  // Run a query without reading the results
  db.run("CREATE TABLE IF NOT EXISTS test (col1, col2);");
  // Insert two rows: (1,111) and (2,222)
  db.run("INSERT INTO test VALUES (?,?), (?,?)", [1,111,2,222]);

  // Prepare a statement
  const stmt = db.prepare("SELECT * FROM test WHERE col1 BETWEEN $start AND $end");
  stmt.getAsObject({$start:1, $end:1}); // {col1:1, col2:111}

  // Bind new values
  stmt.bind({$start:1, $end:2});
  while(stmt.step()) { //
    const row = stmt.getAsObject();
    console.log('Here is a row: ' + JSON.stringify(row));
  }

  const data = db.export();
 // const buffer = new Buffer(data);
 // fs.writeFileSync("filename.sqlite", buffer);
  const a = document.createElement('a');
  const blob = new Blob([data]);
  a.href = URL.createObjectURL(blob);
  a.download = 'database.sqlite';                     //filename to download
  a.click();
}

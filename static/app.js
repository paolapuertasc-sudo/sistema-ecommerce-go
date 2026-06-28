const API = '/api';
let productos = [];
let clientes = [];
let pedidos = [];

const $ = (id) => document.getElementById(id);

function showToast(message, type = 'ok') {
  const toast = $('toast');
  toast.textContent = message;
  toast.className = `toast ${type}`;
  toast.style.display = 'block';
  setTimeout(() => toast.style.display = 'none', 3200);
}

async function request(url, options = {}) {
  const response = await fetch(url, options);
  const data = await response.json();
  if (!response.ok) throw new Error(data.error || 'Ocurrió un error');
  return data;
}

async function cargarTodo() {
  await Promise.all([cargarProductos(), cargarClientes(), cargarPedidos()]);
}

async function cargarProductos() {
  const data = await request(`${API}/productos`);
  productos = data.productos || [];
  $('totalProductos').textContent = productos.length;
  $('productosTabla').innerHTML = productos.map(p => `
    <tr><td>${p.id}</td><td>${p.nombre}</td><td>$${p.precio.toFixed(2)}</td><td>${p.stock}</td></tr>
  `).join('');
  $('pedidoProducto').innerHTML = productos.map(p => `<option value="${p.id}">${p.nombre} · Stock ${p.stock}</option>`).join('');
}

async function cargarClientes() {
  const data = await request(`${API}/clientes`);
  clientes = data.clientes || [];
  $('totalClientes').textContent = clientes.length;
  $('clientesTabla').innerHTML = clientes.map(c => `
    <tr><td>${c.id}</td><td>${c.nombre}</td><td>${c.correo}</td></tr>
  `).join('');
  $('pedidoCliente').innerHTML = clientes.map(c => `<option value="${c.id}">${c.nombre}</option>`).join('');
}

async function cargarPedidos() {
  const data = await request(`${API}/pedidos`);
  pedidos = data.pedidos || [];
  $('totalPedidos').textContent = pedidos.length;
  if (pedidos.length === 0) {
    $('pedidosLista').innerHTML = '<p>No existen pedidos registrados todavía.</p>';
    return;
  }
  $('pedidosLista').innerHTML = pedidos.map(p => `
    <article class="order-card">
      <h3>Pedido #${p.id} · ${p.cliente.nombre}</h3>
      <p>${new Date(p.fecha).toLocaleString()}</p>
      <p>${p.items.map(i => `${i.producto.nombre} x ${i.cantidad}`).join(' · ')}</p>
      <p>Subtotal: $${p.subtotal.toFixed(2)} · IVA: $${p.iva.toFixed(2)}</p>
      <p class="money">Total: $${p.total.toFixed(2)}</p>
    </article>
  `).join('');
}

$('productoForm').addEventListener('submit', async (e) => {
  e.preventDefault();
  try {
    await request(`${API}/productos`, {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify({
        nombre: $('productoNombre').value,
        precio: Number($('productoPrecio').value),
        stock: Number($('productoStock').value)
      })
    });
    e.target.reset();
    showToast('Producto registrado correctamente');
    await cargarProductos();
  } catch (error) { showToast(error.message, 'error'); }
});

$('clienteForm').addEventListener('submit', async (e) => {
  e.preventDefault();
  try {
    await request(`${API}/clientes`, {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify({
        nombre: $('clienteNombre').value,
        correo: $('clienteCorreo').value
      })
    });
    e.target.reset();
    showToast('Cliente registrado correctamente');
    await cargarClientes();
  } catch (error) { showToast(error.message, 'error'); }
});

$('pedidoForm').addEventListener('submit', async (e) => {
  e.preventDefault();
  try {
    await request(`${API}/pedidos`, {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify({
        cliente_id: Number($('pedidoCliente').value),
        items: [{producto_id: Number($('pedidoProducto').value), cantidad: Number($('pedidoCantidad').value)}]
      })
    });
    $('pedidoCantidad').value = '';
    showToast('Pedido creado correctamente');
    await cargarTodo();
  } catch (error) { showToast(error.message, 'error'); }
});

cargarTodo().catch(error => showToast(error.message, 'error'));

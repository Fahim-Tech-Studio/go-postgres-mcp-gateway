import React, { useState } from 'react';

export default function App() {
  const [activeTab, setActiveTab] = useState('overview');

  return (
    <div className="flex h-screen bg-slate-950 text-slate-100">
      {/* Sidebar Navigation */}
      <aside className="w-64 border-r border-slate-800 bg-slate-900/50 p-6 flex flex-col justify-between">
        <div>
          <div className="flex items-center gap-2 mb-8">
            <div className="h-3 w-3 rounded-full bg-cyan-500 animate-pulse"></div>
            <span className="font-bold tracking-wider text-sm text-slate-200">FAHIM TECH STUDIO</span>
          </div>
          <nav className="space-y-2">
            <button 
              onClick={() => setActiveTab('overview')}
              className={`w-full text-left px-4 py-2.5 rounded-lg text-sm font-medium transition-colors ${activeTab === 'overview' ? 'bg-cyan-500/10 text-cyan-400 border border-cyan-500/20' : 'text-slate-400 hover:text-slate-200 hover:bg-slate-800/50'}`}>
              📊 Cluster Overview
            </button>
            <button 
              onClick={() => setActiveTab('guardrails')}
              className={`w-full text-left px-4 py-2.5 rounded-lg text-sm font-medium transition-colors ${activeTab === 'guardrails' ? 'bg-cyan-500/10 text-cyan-400 border border-cyan-500/20' : 'text-slate-400 hover:text-slate-200 hover:bg-slate-800/50'}`}>
              🛡️ SQL AST Guardrails
            </button>
            <button 
              onClick={() => setActiveTab('audit')}
              className={`w-full text-left px-4 py-2.5 rounded-lg text-sm font-medium transition-colors ${activeTab === 'audit' ? 'bg-cyan-500/10 text-cyan-400 border border-cyan-500/20' : 'text-slate-400 hover:text-slate-200 hover:bg-slate-800/50'}`}>
              🔒 PII Masking Logs
            </button>
          </nav>
        </div>
        <div className="text-xs text-slate-500 border-t border-slate-800/80 pt-4">
          Go PostgreSQL MCP Gateway v1.0
        </div>
      </aside>

      {/* Main Content Area */}
      <main className="flex-1 overflow-y-auto p-8">
        <header className="flex justify-between items-center mb-8 pb-4 border-b border-slate-800">
          <h1 className="text-xl font-semibold capitalize text-slate-100">{activeTab} Control Center</h1>
          <span className="px-3 py-1 bg-emerald-500/10 border border-emerald-500/20 text-emerald-400 rounded-full text-xs font-mono">
            ● System Healthy
          </span>
        </header>

        {activeTab === 'overview' && (
          <div className="grid grid-cols-3 gap-6">
            <div className="p-6 rounded-xl bg-slate-900 border border-slate-800">
              <h3 className="text-slate-400 text-sm font-medium mb-1">Active Connections</h3>
              <p className="text-3xl font-bold text-cyan-400 font-mono">12</p>
            </div>
            <div className="p-6 rounded-xl bg-slate-900 border border-slate-800">
              <h3 className="text-slate-400 text-sm font-medium mb-1">Blocked Mutations</h3>
              <p className="text-3xl font-bold text-emerald-400 font-mono">148</p>
            </div>
            <div className="p-6 rounded-xl bg-slate-900 border border-slate-800">
              <h3 className="text-slate-400 text-sm font-medium mb-1">PII Columns Masked</h3>
              <p className="text-3xl font-bold text-purple-400 font-mono">1,230</p>
            </div>
          </div>
        )}

        {activeTab === 'guardrails' && (
          <div className="p-6 rounded-xl bg-slate-900 border border-slate-800">
            <h3 className="text-lg font-medium mb-4">AST Security Rule Enforcement</h3>
            <p className="text-sm text-slate-400 mb-4">All incoming model queries are parsed to ensure strict compliance with read-only boundaries.</p>
            <div className="font-mono text-xs bg-slate-950 p-4 rounded-lg border border-slate-800 text-emerald-400">
              [ALLOWLIST ACTIVE] SELECT, WITH statements permitted.<br/>
              [BLOCKLIST ACTIVE] DROP, INSERT, UPDATE, DELETE, ALTER, TRUNCATE.
            </div>
          </div>
        )}

        {activeTab === 'audit' && (
          <div className="p-6 rounded-xl bg-slate-900 border border-slate-800">
            <h3 className="text-lg font-medium mb-4">Live PII Column Scrubbing Audit Trail</h3>
            <div className="space-y-3 font-mono text-xs">
              <div className="p-3 bg-slate-950 rounded border border-slate-800 flex justify-between items-center">
                <span className="text-slate-300">Query: SELECT * FROM users LIMIT 5;</span>
                <span className="text-purple-400">Masked: password_hash, api_key</span>
              </div>
            </div>
          </div>
        )}
      </main>
    </div>
  );
}
